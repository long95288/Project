package main

import "fmt"

func main() {
	const (
		a = iota // 0
		b	// 1
		c 	//2
		d = "ha" // iota +=1 3
		e // ha 4
		f = 100 //5
		g //6
		h = iota //7 恢复计数
		i // 8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
}
