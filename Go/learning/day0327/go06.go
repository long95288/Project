package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum)/float32(count)
	fmt.Printf("mean的值为：%f \n",mean)

	var num int = 10
	var sub int = 3
	fmt.Printf("10(int) /3(int) = %d \n",num/sub)
	// 不允许 10(int) / 3(float32)
	fmt.Printf("10(int) /3(float32) = %f \n",(float32(num) / float32(sub)))
}
