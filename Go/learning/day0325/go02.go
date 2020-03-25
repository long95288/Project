package main

import "fmt"

func main() {
	x,y := "Google","Golang"

	a,b := swap(x,y)
	fmt.Println(a,b)
	a,b = addHead(x,y)
	fmt.Println(a,b)
	fmt.Println(x,y)
	c,d := 100,200
	// 将指针地址传入交换函数
	swap2(&c,&d)
	fmt.Println(c,d)

}

func swap(x, y string) (string, string) {
	return y,x
}
// 按值引用
func addHead(x, y string) (string, string) {
	x = "Head-"+x
	y = "Head-"+y
	return x,y
}
/**
引用传递
 */
 func swap2(x *int,y *int){
 	var temp int
 	temp = *x
 	*x = *y
 	*y = temp
 }
