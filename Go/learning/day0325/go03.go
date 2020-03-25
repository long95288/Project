package main

import "fmt"

/**
声明为全局变量
 */
var g int

var g2 int = 20
func main() {
	// 声明为局部变量
	var a,b int

	// 初始化参数
	a = 10
	b = 20
	g = a + b
	// 局部变量优先被考虑
	var g2 string = "Hello World"
	fmt.Println(g)
	fmt.Println(g2)

}
