package main

func test(f func())  {
    f()
}
func test2() func(int,int) int{
    return func(x,y int) int {
        return x + y
    }
}

func main() {
    // 直接执行匿名函数
    func(s string) {
        println(s)
    }("Hello World")
    // 赋值给变量
    add := func(x,y int) int{
        return x+y
    }
    println(add(1,2))
    // 作为参数
    test(func() {
        println("Hello World")
    })
    // 作为返回值
    test2 := test2()
    println(test2(1,2))
}
