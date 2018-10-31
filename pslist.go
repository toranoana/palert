package main

import (
    "fmt"
    "github.com/mitchellh/go-ps"
)

func main() {

    processes, _ := ps.Processes()

    for i, p := range processes {
        fmt.Printf("%d : %s\n", i, p.Executable())
    }
}

