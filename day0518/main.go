package main

import "fmt"

func main() {
	str := "Hello被"
	fmt.Println("str len = ", len(str))

	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符串=%c", r[i])
	}

}
