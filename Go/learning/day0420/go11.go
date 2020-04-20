package main

import (
    "log"
    "os"
)

func main() {
    f,err := os.Open("day0420/go01.go")
    if err != nil {
        log.Fatalln(err)
    }
    // 执行结束后关闭文件
    defer f.Close()
    
    var data []byte
    f.Read(data)
    println(data)
}
