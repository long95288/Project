package main

import "fmt"

func main() {
	var a int = 10
	var s = "String"
	fmt.Println("变量的地址为：%x\n",&a)
	// 声明一个指针变量
	// 整型指针变量
	var ap *int
	// 字符串指针变量
	var sp *string = &s

	// 赋值操作
	ap = &a
	*ap = 20
	*sp = "Hello World"
	fmt.Println(a)
	fmt.Println(s)

	var nullPtr *int
	if nullPtr == nil{
		fmt.Println("空指针")
		fmt.Println(nullPtr)
	}else {
		fmt.Println("非空指针")
	}
	nullPtr = &a
	fmt.Println(nullPtr)
	fmt.Println(&a)

}
