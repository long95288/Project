package main

func main() {
    x,y := 1,2
    
    defer func(a int) {
        println("Defer x,y = ",a,y)
    }(x)
    
    x += 100
    y += 200
    println(x,y)
}
