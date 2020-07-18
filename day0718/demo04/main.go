package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    defer trace("bigSlowOperation")()
    time.Sleep(3*time.Second)
    _ = double(3)
    fmt.Println(triple(4))
    
}
func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s",msg)
    return func() {
        log.Printf("exit %s (%s)",msg,time.Since(start))
    }
}
// 在defer函数中查看返回值的内容
//
func double(x int) (result int){
    defer func() {fmt.Printf("double(%d) = %d \n",x,result)}()
    return x + x
}
// 修改返回值给调用者
func triple(x int) (result int) {
    defer func() {result += x}()
    return double(x)
}
