package m01

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为:  %d", area)

	println()
	println(a, b, c)

}