package main

func main() {
	a := struct {
		x int
	}{}

	a.x = 100
	p := &a
	p.x += 100

	println(p.x)
}
