package main
func test(x byte){
	println(x)
}
func main(){
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	// byte 是 uint8的别名
	test(c)
	var x int = 100
	//必须显示转换
	var y int64 = int64(x)
	// int 和 int64必须显式转换
	x = add(x,int(y))
	println(x)
}

func add(x,y int) int{
	return x + y
}
