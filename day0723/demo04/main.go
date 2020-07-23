package main

import (
    "fmt"
    "reflect"
)
// kind处理不同分支
func printType(t reflect.Type) {
    switch t.Kind() {
    case reflect.Int:
        fmt.Println("int")
    case reflect.String:
        fmt.Println("string")
    }
}
func main() {
    a := 1
    t := reflect.TypeOf(a)
    printType(t)
    t = reflect.TypeOf("Hello World")
    printType(t)
}
