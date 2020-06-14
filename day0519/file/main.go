package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	defer file.Close()
	if nil != err {
		fmt.Println("打开文件失败")
	}
	fmt.Printf("flie = %v\n", file)
	var b []byte = make([]byte, 1024*1024)
	var data []byte
	n, err := file.Read(b)
	// for {
	// 	if n == -1 || err != nil {
	// 		break
	// 	}
	// }
	fmt.Println("n = ", n, "err = ", err)
	data = b[:n]
	fmt.Printf("%q \n", string(data))

	// reader := bufio.NewReader(file)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
}
