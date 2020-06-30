package main

import (
    "fmt"
    "os"
)

func main() {
    dir := "download/"
    file,err := os.Open(dir)
    if err != nil {
        fmt.Println(err)
        return
    }
}
