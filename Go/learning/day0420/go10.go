package main

func test(x int) (func(), func()) {
    return func() {
        println(x)
        x += 10
    }, func() {
            println(x)
        }
}
func main() {
    a,b := test(100)
    // 改变变量
    a()
    // 获得变量
    b()
}
