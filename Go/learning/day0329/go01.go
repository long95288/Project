package main

func consumer(data chan int, done chan bool) {
	for x := range data{
		println("recv:",x)
	}
	done <- true
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	// 关闭通道
	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)
	//
	println("进程阻塞")
	// 一直等待，知道done中含有数据
	<-done
	println("阻塞结束")
}
