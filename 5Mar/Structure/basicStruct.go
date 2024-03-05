package main

import "fmt"

type Book struct {
	title string 
	author string
	subject string
	book_id int
}

func main(){
	var book1 Book
	book1.title = "Understanding Physics"
	book1.author = "H C Verma"
	book1.subject = "Physics"
	book1.book_id = 12345

	fmt.Printf("%s\n", book1.title)
	fmt.Printf("%s\n", book1.author)
	fmt.Printf("%s\n", book1.subject)
	fmt.Printf("%d\n", book1.book_id)
}