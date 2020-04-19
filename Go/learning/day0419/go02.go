package main

import "fmt"

func main() {
    a := 1.0 << 3
    fmt.Printf("%T, %v \n",a,a)
    
    var s uint = 3
    // 会先将1.0转成float32类型,然后位移出错
    b := 1.0 << s
    fmt.Printf("%T,%v \n",b,b)
    
    var c int32 = 1.0 << s
    fmt.Printf("%T, %v \n",c,c)
}
