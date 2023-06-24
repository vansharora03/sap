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
        fmt.Println("    create <name>    :    create a new sap script")
        fmt.Println("")

    } else if args[0] == "init" {
        // init
        if err := sap_init(); err != nil {
            fmt.Println(err)
        }

    } else if args[0] == "create" {
        // create

        // Do not do anything if sap.json does not exist
        if _, err := os.Stat("sap.json"); os.IsNotExist(err) {
            fmt.Println("sap.json file not detected")
            fmt.Println("Please initialize a sap.json file with sap init")
        }

    }


}
