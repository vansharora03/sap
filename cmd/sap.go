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
        fmt.Println("    list    :    list all current scripts with their commands")
        fmt.Println("")

    } else if args[0] == "init" {
        // init
        if err := sap_init(); err != nil {
            fmt.Println(err)
        }

    } else if args[0] == "create" {
        // create
        // Check if name was given
        if len(args) < 2 {
            fmt.Println("Please give a name for the script i.e sap create foo")
            return
        }

        if err := sap_create(args[1]); err != nil {
            fmt.Println(err)
        }

    } else if args[0] == "list" {
        //list
        if err := sap_list(); err != nil {
            fmt.Println(err)
        }
    }


}
