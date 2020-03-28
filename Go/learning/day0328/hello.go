package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name")
	// 以\n为分隔符读取一段内容
	input,err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("Found an error: %s \n",err)
	}else{
		// 去掉最后一个字符\n
		input = input[:len(input)-1]
		fmt.Printf("Hello,%s \n",input)
	}
}
