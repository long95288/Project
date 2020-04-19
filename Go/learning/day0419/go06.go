package main

func main() {
    switch x := 5; x {
    case 5:
        x += 10
        println(x)
        fallthrough
    case 6:
        x += 20
        println(x)
    default:
        println(x)
    }
}
