package main

import "fmt"

func main() {
	var number []int
	fmt.Printf("len=%d cap=%d slice=%v \n",len(number),cap(number),number)

	number = append(number,1)
	printSlice(number)

	number = append(number,1,2,3,4,5,6)
	printSlice(number)

	numbers2 := make([]int,len(number),(cap(number)*2))
	printSlice(numbers2)

	/**
	拷贝数据
	 */
	copy(numbers2,number)
	printSlice(numbers2)

}
func printSlice(slice []int)  {
	fmt.Printf("len=%d cap=%d slice=%v \n",len(slice),cap(slice),slice)
}
