package main

import "fmt"

func main() {
	fmt.Printf("max 100,200 is %d",max(100,200))
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
