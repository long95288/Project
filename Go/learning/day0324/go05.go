package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	if a == b {
		fmt.Printf("a == b \n")
	}else {
		fmt.Printf("a != b \n")
	}

	if a < b {
		fmt.Printf("a<b \n")
	}else {
		fmt.Printf("a>b \n")
	}

	if a > b {
		fmt.Printf("a>b \n")
	}else {
		fmt.Printf("a<b \n")
	}

	a = 5
	b = 20
	if a < b{
		fmt.Printf("a<b \n")
	}else {
		fmt.Printf("a>b \n")
	}
}
