package main

import "fmt"

func main() {
	numbers := []int{0,1,2,3,4,5,7,8,9}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)

	// [1,4) 从1开始,到4-1结束
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	fmt.Println("numbers[:3] ==",numbers[:3])

	fmt.Println("numbers[2:0] ==",numbers[2:])

	numbers2 := numbers[1:4]

	fmt.Printf("numbers2 len=%d cap=%d slice=%v \n",len(numbers2),cap(numbers2),numbers2)

}
