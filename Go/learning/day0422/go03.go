package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    s := "abcdefg"
    
    s1 := s[:3]
    s2 := s[1:4]
    s3 := s[2:]
    
    println(s1,s2,s3)
    
    fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s)))
    fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s1)))
    
    
}
