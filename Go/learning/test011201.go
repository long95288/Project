package main

import "fmt"

func main() {
	const LEGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LEGTH * WIDTH
	fmt.Println("面积为:  %d", area)

	println()
	println(a, b, c)

}
