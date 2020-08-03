package demo01

import (
    "fmt"
    "testing"
)
// 测试函数
func TestHelloWorld(t *testing.T) {
    fmt.Println("Function Test Hello world")
}

func BenchmarkHello(t *testing.B)  {
    fmt.Println("Benchmark Test Hello")
   
}

func ExampleHello()  {
    //fmt.Println(false)
    fmt.Println(true)
    // Output:
    // true
    // false
}