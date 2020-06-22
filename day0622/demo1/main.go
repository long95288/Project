package main

import (
    "fmt"
    "os"
)

func main() {
    pid := os.Getpid()
    ppid := os.Getppid()
    fmt.Println("pid = ",pid)
    fmt.Println("ppid = ",ppid)
}
