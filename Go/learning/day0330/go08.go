package main

import "fmt"

func main() {
	type data [2]int
	// 基础类型相同，右值为未命名类型
	var d data = [2]int{1,2}
	fmt.Println(d)
	a := make(chan int, 2)
	// 双向通道转为单向通道
	var b chan <- int = a
	b <- 2
}
