package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main() {

	var b Books
	b.author= "Hello World"
	b.title="First"
	b.book_id=123
	fmt.Println(b)

	fmt.Println(Books{title:"Hello",author:"JJ"})
	fmt.Println(Books{"Hello","","",1234})

	title := b.title
	println(title)
	printBook(b)

}
func printBook(book Books) {
	fmt.Println(book)
}

