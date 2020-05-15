package main

import "fmt"

func main() {
    type (
        user struct {
            name string
            age uint8
        }
        event func(string) bool
        )
    
    u := user{
        name: "Tom",
        age:  20,
    }
    fmt.Println(u)
    
    var f event = func(s string) bool {
        println(s)
        return s != ""
    }
    f("abc")
}
