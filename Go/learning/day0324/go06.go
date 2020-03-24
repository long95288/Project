package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	if a&&b{
		fmt.Println("True")
	}else {
		fmt.Println("False")
	}

	if a||b {
		fmt.Println("True")
	}else {
		fmt.Println("False")
	}

	a = false
	b = true
	if !(a && b) {
		fmt.Println("True")
	}
}
