package main

func main() {
	var a,b struct{}

	println(&a,&b)
	println(&a == &b,&a == nil)
}
