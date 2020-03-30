package main

import "fmt"

func main() {
	type (
		// 结构体
		user struct {
			name string
			age uint8
		}
		// 函数类型
		event func(string) bool
	)
	u := user{"TOM",20}
	fmt.Println(u)

	// 实现类
	var f event = func(s string) bool {
		println(s)
		return s != ""
	}
	f("abc")
}
