package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args[1:]

    if len(args) == 0 {
        // Show help menu on no arguments
        fmt.Println("")
        fmt.Println("sap is a tool for writing and storing scripts")
        fmt.Println("")
        fmt.Println("Version:")
        fmt.Println("")
        fmt.Println("    sap version 0")
        fmt.Println("")
        fmt.Println("Usage:")
        fmt.Println("")
        fmt.Println("    sap <command> [arguments]")
        fmt.Println("")
        fmt.Println("Commands:")
        fmt.Println("")
        fmt.Println("    init    :    start sap in a new project")
        fmt.Println("")

    } else if args[0] == "init" {
        // init

        // Do not overwrite an existing sap.json
        if _, err := os.Stat("sap.json"); !os.IsNotExist(err) {
            fmt.Println("sap.json file already exists.")
            return
        }

        _, err := os.Create("sap.json")
        if err != nil {
            fmt.Print(err.Error());
        }

    }


}
