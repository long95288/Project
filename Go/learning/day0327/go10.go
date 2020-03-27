package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _,v := range s{
		sum += v
	}
	// sum 发送到c
	c <- sum
}
func main() {
	s := []int{7,2,8,-9,4,0}
	// 声明一个通道
	c := make(chan int)
	go sum(s[:len(s)/2],c) // 运行结果17
	go sum(s[len(s)/2:],c) // 运行结果-5
	// 接收通道的c
	x,y := <-c,<-c

	fmt.Println(x,y,x+y)

}
