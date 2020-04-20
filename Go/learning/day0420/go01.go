package main

func hello(){
    println("Hello, world")
}
func getHello() func() {
    return hello
}
func exec(f func())  {
    f()
}
func main() {
    f := hello
    exec(f)
    f = getHello()
    exec(f)
}
