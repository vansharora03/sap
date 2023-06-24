package main
import (
    "os"
    "fmt"
)

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


// func sap_create(name string) error {}
