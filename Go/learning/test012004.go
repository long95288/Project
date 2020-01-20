package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var re int
	re = max(a, b)
	fmt.Println("最大值: %d", re)

	c, d := swap("Hello", "World")
	println(c, d)

}
func max(num1, num2 int) int {
	/*局部变量*/
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/**
交换值
*/
func swap(x, y string) (string, string) {
	return y, x
}
