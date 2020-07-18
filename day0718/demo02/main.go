package main

import "fmt"

func sum(vals... int) int{
    total := 0
    for _,val := range vals{
        total += val
    }
    return total
}

func main() {
    fmt.Println(sum())
    fmt.Println(sum(3))
    fmt.Println(sum(1,2,3,4,5))
    values := []int{1,2,3,4,5}
    // 使用数组或者切片的时候需要后面添加3个.
    fmt.Println(sum(values...))
    
}
