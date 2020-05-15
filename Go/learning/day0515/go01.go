package main

import "fmt"

func main() {
    var name string
    var age byte
    var sal float32
    var isPass bool

    fmt.Println("姓名 年龄 薪水 是否通过考试")
    fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
    fmt.Printf("姓名 %s 年龄%d 薪水%f ispass %t",name,age,sal,isPass)
}
