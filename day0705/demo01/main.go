package main

import "fmt"

func sum(values ...int) int {
    total := 0
    for _,v := range values {
        total += v
    }
    return total
}

func main() {
    re := sum(1,2,3,4,5)
    fmt.Println(re)
    re = sum(3,3,3)
    fmt.Println(re)
}
