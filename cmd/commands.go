package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// A script is an executable that has a name and a body of
// other executables.
type script struct {
    Name string `json:"name"`
    Body []executable `json:"body"`
}

// An executable is an individual command within a script.
// This can either be a normal command, or a call to another script.
type executable interface {
    Exec() (string, error)
}


// fileExists returns whether the fileName exists.
func fileExists(fileName string) bool {
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        return false
    }
    return true
}


// sap_init performs the functions of sap init.
// Creates a blank sap.json file, only when one
// does not exist already.
func sap_init() error {
    // Do not overwrite an existing sap.json
    if fileExists("sap.json") {
        fmt.Println("sap.json file already exists.")
        return nil
    }

    _, err := os.Create("sap.json")
    if err != nil {
        return err
    }

    fmt.Println("Successfully created sap.json")

    return nil
}


// sap_create creates a new sap script with the given name
func sap_create(name string) error {
    // Do not do anything if sap.json does not exist
    if !fileExists("sap.json") {
        fmt.Println("No sap.json detected. Please execute sap init")
        return nil
    }

    // Read file into contents
    contents, err := os.ReadFile("sap.json")
    if err != nil {
        return err
    }
    
    // Get the array of scripts
    var scripts []script
    json.Unmarshal(contents, &scripts)
    for _, script := range scripts {
        if script.Name == name {
            return errors.New("This script name already exists in sap.json")
        }
    }

    // Update scripts
    scr := script{Name: name}
    scr.Body = []executable{}
    scripts = append(scripts, scr)

    // Write updated scripts to the file
    js, err := json.MarshalIndent(scripts, "", "\t")
    if err != nil {
        return nil
    }
    err = os.WriteFile("sap.json", js, fs.ModeAppend)
    if err != nil {
        return err
    }

    return nil
}
