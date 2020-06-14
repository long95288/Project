package main

import (
	"fmt"
)

// Cat 猫的结构体
type Cat struct {
	Name string
	Age  int
}

func test1() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m2 := make(map[string]string, 10)

	m1["c1"] = "beijing"
	m1["c2"] = "tianjin"

	m2["h1"] = "text1"
	m2["h2"] = "text2"

	mapChan <- m1
	mapChan <- m2

	fmt.Printf("mapChanel %v \n", mapChan)
	fmt.Printf("mapChanel 1 %v \n", <-mapChan)
	fmt.Printf("mapChanel 2 %v \n", <-mapChan)

}

func test2() {
	var catChan chan Cat
	var catChan2 chan *Cat
	catChan = make(chan Cat, 10)
	catChan2 = make(chan *Cat, 10)
	cat1 := Cat{
		Name: "tom",
		Age:  12,
	}
	cat2 := Cat{
		"tom2", 32,
	}
	catChan <- cat1
	catChan <- cat2
	catChan2 <- &cat1
	catChan2 <- &cat2
	fmt.Printf("cat1 = %v cat2 = %v \n", <-catChan, <-catChan)
	fmt.Printf("cat1 = %v cat2 = %v \n", *<-catChan2, <-catChan2)

}
func test3() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan <- Cat{"tom2", 22}
	allChan <- 32
	allChan <- "Hello Word"
	allChan <- map[string]int{"h1": 1, "h2": 2}
	allChan <- []int{1, 2, 3, 4}
	close(allChan)
	for i := range allChan {
		fmt.Printf("i = %v\n", i)
	}
}
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// time.Sleep(time.Second * 1)
		intChan <- i
		fmt.Println("WriteData", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		// time.Sleep(time.Second * 1)
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到的数据: %v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

func test4() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// time.Sleep(time.Second * 3)
	fmt.Println("main process start")
	for {
		fmt.Println("waiting end1!")
		_, ok := <-exitChan
		if !ok {
			break
		}
		fmt.Println("waiting end2!")

	}
	fmt.Println("main process exit")
}
func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值=%v intChan 地址=%p\n", intChan, &intChan)
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	// intChan <- 98
	// 赋值操出
	fmt.Printf("channel len=%v cap=%v \n ", len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len=%v cap=%v \n ", len(intChan), cap(intChan))

	num3 := <-intChan
	num4 := <-intChan
	num5 := 0
	fmt.Printf("num3 = %v num4 = %v num5 = %v \n ", num3, num4, num5)

	test1()
	test2()
	test3()
	test4()
}
