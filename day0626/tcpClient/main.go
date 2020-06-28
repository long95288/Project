package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    conn,err := net.Dial("tcp",":8080")
    if err != nil {
        fmt.Println("err :",conn)
        return
    }
    defer conn.Close()
    // 控制台输入
    r := bufio.NewReader(os.Stdin)
    for{
        input,err := r.ReadString('\n')
        inputInfo := strings.Trim(input,"\r\n")
        if err != nil {
            fmt.Println("input err :",err)
            break
        }
        //input += "\n"
        _,err = conn.Write([]byte(inputInfo))
        if err != nil {
            fmt.Println("send err ,",err)
            return
        }
        buf := [512]byte{}
        n,err :=conn.Read(buf[:])
        if err != nil {
            return
        }
        fmt.Println(string(buf[:n]))
    }
}
