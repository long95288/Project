package main

import (
	"fmt"
	"runtime"
)
// 先初始化后运行main
func init()  {
	fmt.Printf("Map: %v \n",m)
	info = fmt.Sprintf("OS: %s, Arch: %s",runtime.GOOS,runtime.GOARCH)
}

var m = map[int]string{1:"A",2:"B",3:"C"}

var info string

func main() {
	fmt.Println(info)
}
