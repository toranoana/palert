package main

import (
    "fmt"
    "github.com/mitchellh/go-ps"
)

func main() {

    processes, _ := ps.Processes()

    for _, p := range processes {
        fmt.Printf("pid=%d ppid=%d %s\n", p.Pid(), p.PPid(), p.Executable())
    }
}

