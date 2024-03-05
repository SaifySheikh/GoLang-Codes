package main

import "fmt"

type Books struct {
	title string 
	author string
	subject string
	book_id int
}

func main(){
	var book1 Books
	var book2 Books
	book1.title = "Understanding Physics"
	book1.author = "H C Verma"
	book1.subject = "Physics"
	book1.book_id = 12345

	book2.title = "Understanding Mathematics"
	book2.author = "A Agrawal"
	book2.subject = "Mathematics"
	book2.book_id = 123456

	printBook(book1)
	printBook(book2)
}


func printBook(book Books){
	fmt.Printf("%s\n", book.title)
	fmt.Printf("%s\n", book.author)
	fmt.Printf("%s\n", book.subject)
	fmt.Printf("%d\n", book.book_id)
	fmt.Printf("\n")

}