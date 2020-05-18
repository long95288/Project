package main

import (
    "fmt"
    "github.com/pkg/errors"
)

func test() {
    defer func() {
        err := recover()
        if err!= nil{
            fmt.Println("err = ",err)
        }
    }()
    
    num1 := 10
    num2 := 0
    res := num1/num2
    fmt.Println("res = ",res)
}

func readConf(name string) (err error) {
    if name == "conf.ini" {
        return nil
    }
    return errors.New("error file")
}
func test02() {
    err := readConf("conf.ini")
    if err != nil{
        panic(err)
    }
    fmt.Println(".........")
}
func main() {
    test()
    fmt.Println(".....")
    test02()
    fmt.Println("...........")
}
