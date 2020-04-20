package main

func main() {
    defer println("a")
    defer println("b")
}
