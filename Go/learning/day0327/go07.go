package main

import "fmt"

/**
定义一个接口Phone
函数有call
*/
type Phone interface {
	call()
}
type NokiaPhone struct {

}
// 实现接口
func (nokiaPhone NokiaPhone) call(){
	fmt.Println("I am Nokia,I can call you")
}

type IPhone struct {

}

// 实现call接口
func (iPhone IPhone) call()  {
	fmt.Println("I am IPhone,I can call you")
}
func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
