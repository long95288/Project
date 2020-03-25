package main

import "fmt"

func main() {
	const MAX int = 3
	a:=[]int{1,2,3}
	// *int为数组的类型,里面存放着数组指针
	var ptr [MAX]*int
	for i:=0;i<MAX;i++{
		// 指针赋值
		ptr[i] = &a[i]
	}
	for i:=0;i<MAX;i++{
		fmt.Printf("a[%d] = %d\n",i,*ptr[i])
	}
}
