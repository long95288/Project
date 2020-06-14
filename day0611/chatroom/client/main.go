package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	var key int
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统------------------------")
		fmt.Println("\t\t\t 1 登入聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d \n", key)
		switch key {
		case 1:

			loop = false
		case 2:
			loop = false
		case 3:
			loop = true
			os.Exit(0)
		default:
			fmt.Println("错误选项")
		}

	}
}
