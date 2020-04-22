package main

import "fmt"

func main() {
    s := "ab"+"cd"
    println(s == "abcd")
    println(s > "abc")
    
    fmt.Printf("%d \n",s[1])
    // 不允许获得元素地址
    // println(&s[1])
}
