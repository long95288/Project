package main

import "fmt"

func test(a ...int){
    // fmt.Println(a)
    for i:=range a {
        a[i] += 100
    }
}

func main() {
    a :=[]int{10,20,30}
    test(a...)
    fmt.Println(a)
}
