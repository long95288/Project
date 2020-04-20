package main

func test(p **int){
    x := 100
    *p = &x
}
func main() {
    var p *int
    test(&p)
    println(*p)
}
