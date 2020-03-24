package main

import "fmt"

const (
	i=1<<iota
	j=3<<iota
	k // 3左移2位 12 解:11 << 2 => 1100 ->12
	l // 3左移3位 24 解:11 << 3 => 11000 -> 24
)

func main() {
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
}
