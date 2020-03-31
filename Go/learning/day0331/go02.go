package main

func main() {
	x := 10

	var p *int = &x
	*p += 20

	println(p,*p)
}
