package main

func test() *int {
    a := 0x100
    // 返回局部变量指针
    return &a
}

func main() {
    var a *int = test()
    println(a,*a)
}
