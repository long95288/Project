package main

import "fmt"

// flags 为 byte别名
type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	f := read | exec
	fmt.Printf("%b \n",f)
}
