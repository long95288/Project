package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello Word")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test panic ", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}
