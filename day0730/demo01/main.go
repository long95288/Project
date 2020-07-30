package main

import (
    "errors"
    "fmt"
)

func TestPanic() {
    //defer func() {
    //   if err := recover();err!= nil{
    //       fmt.Println("Function Catch Panic ",err)
    //   }
    //}()
    // 抛出panic
    panic(errors.New("Test Panic"))
}
func main() {
   // TestPanic() //无法捕获
    // defer必须在TestPanic之前才能捕获
    defer func() {
        if err := recover();err!= nil{
            fmt.Println("Main catch Panic ",err)
        }
    }()
    TestPanic()
    fmt.Println("over")
}
