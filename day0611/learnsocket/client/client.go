package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if nil != err {
		fmt.Println("client dial err = ", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if nil != err {
			fmt.Println("readString err = ", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}
		_, err = conn.Write([]byte(line))
		if nil != err {
			fmt.Println("conn.Write err = ", err)
		}
	}
}
