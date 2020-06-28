package main

import (
    "bufio"
    "fmt"
    "net"
)

func process(conn net.Conn){
    defer conn.Close()
    for{
        reader := bufio.NewReader(conn)
        var buf [128]byte
        n,err := reader.Read(buf[:])
        if err != nil {
            fmt.Println("read client err,",err)
            return
        }
        restr := string(buf[:n])
        fmt.Printf("收到的数据:%q",restr)
        conn.Write([]byte("return "+restr))
    }
}

func main() {
    listen,err := net.Listen("tcp",":8080")
    if err != nil {
        fmt.Println("listen err :",err)
        return
    }
    for true {
        conn,err := listen.Accept()
        if err != nil {
            fmt.Println("accept failed ,err ",err)
            continue
        }
        fmt.Println("新连接")
        go process(conn)
    }
}
