package main

func main() {
	x := 100
	println(&x)
	// x为退化赋值,y为新的变量定义
	x,y := 200,"abc"

	println(&x,x)
	println(y)
	{
		x,y := 200,300
		println(&x,x,y)
	}
}
