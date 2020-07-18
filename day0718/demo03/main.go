package main

import "fmt"

func tt(val... int){
    if len(val) < 3{
        fmt.Println("tt 参数数小于3")
    }else{
        fmt.Println("tt 参数的数量为",len(val))
    }
}


func main() {
    tt(1)
    tt(1,2)
    tt(1,2,3)
    tt(1,2,3,4,5,6)
}
