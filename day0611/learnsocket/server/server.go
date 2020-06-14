package main

import (
	"fmt"
	_ "io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 创建切片
		buf := make([]byte, 1024)
		fmt.Printf("服务器等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if nil != err {
			fmt.Printf("客户端退出 err= %v ", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器正在监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if nil != err {
		fmt.Printf("listen err = %v", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Printf("等待客户端连接....")
		conn, err := listen.Accept()
		if nil != err {
			fmt.Println("Accept() err = ", err)
		} else {
			fmt.Printf("Accept() suc con = %v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}

}
