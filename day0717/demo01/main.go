package main

import "fmt"

func main() {
    //i,j := 12,12
    var Array1 [][]int
    Array1 = make([][]int,12)
    //Array1[1] = make([]int,3)
    Array1[2] = append(Array1[2],1)
    Array1[2] = append(Array1[2],2)
    Array1[1] = make([]int,12)
    fmt.Println(Array1)
    Array1[2][1] = 1
    fmt.Println(Array1)
    var arr2 [][]int
    arr2 = make([][]int,12)
    for i:=0;i<len(arr2);i++ {
        arr2[i] = make([]int,5)
    }
    fmt.Println(arr2)
}
