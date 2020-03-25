package main

import "fmt"

func main() {
	// 初始化数组
	var intArray1 [10] int
	var intArray2 = [10]int{1,2,3}
	var intArray3 = [...]int{1,2,3,4}

	fmt.Println(intArray1[9])
	fmt.Println(intArray2[1])
	fmt.Println(intArray3[2])
	intArray3[1] = 0
	fmt.Println(intArray3[1])

}
