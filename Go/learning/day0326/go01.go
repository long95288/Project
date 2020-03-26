package main

import "fmt"

func main() {
	// 3为初始容量 5为切片容量
	var numbers = make([]int,3,5)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)

	var numbers2 []int
	if numbers2 == nil {
		fmt.Printf("切片是空的 \n")
		fmt.Printf("len=%d cap=%d slice=%v \n",len(numbers2),cap(numbers2),numbers2)
	}

}
