package main

import "fmt"

func main() {
    var intArr [5]int = [...]int{1,2,3,4,5}
    
    slice := intArr[1:3]
    fmt.Println("intArr=",intArr)
    fmt.Println("slice = ",slice)
    fmt.Println("slice len = ",len(slice))
    fmt.Println("slice cap = ",cap(slice))
    slice = make([]int,11,11)
    slice[10] = 11
    fmt.Println("slice = ",slice)
    
    slice = append(slice,333,222)
    fmt.Println("slice = ",slice)
    fmt.Println("slice len = ",len(slice))
    fmt.Println("slice cap = ",cap(slice))
    slice = append(slice,slice...)
    fmt.Println("slice = ",slice)
    fmt.Println("slice len = ",len(slice))
    fmt.Println("slice cap = ",cap(slice))
    
    var slice3 = []int{1,2,3,4,5}
    var slice4 = make([]int,10)
    copy(slice4,slice3)
    fmt.Println("slice3 = ",slice3)
    fmt.Println("slice4 = ",slice4)
    
    var slice5 = make([]int,1)
    copy(slice5,slice4)
    fmt.Println("slice5 = ",slice5)
    
    str := "Hello world 汉字"
    slice6 := str[:]
    fmt.Println("slice6",slice6)
    fmt.Printf("slice6[1] %c \n",slice6[1])
    fmt.Printf("slice6 type %T\n",slice6)
    slice7 := []rune(str)
    slice7[1] = '中'
    str = string(slice7)
    fmt.Println("slice7",slice7)
    fmt.Println("str = ",str)
    
}
