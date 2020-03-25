package main

import "fmt"

func main() {
	// 指针的指针
	var pptr **int
	var ptr *int
	var a int = 10
	ptr = &a
	pptr = &ptr
	fmt.Printf("pptr %x\n",pptr)
	fmt.Printf("*pptr %x\n",*pptr)
	fmt.Printf("**pptr %x\n",**pptr)

	fmt.Printf("ptr %x\n",ptr)
	fmt.Printf("*ptr %x\n",*ptr)
	//fmt.Printf("**ptr %x\n",**ptr)
	fmt.Printf("a %d\n",a)


}
