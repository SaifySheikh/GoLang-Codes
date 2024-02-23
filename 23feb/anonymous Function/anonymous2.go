package main

import "fmt"

func main() {
	var greet=func(){
		fmt.Println("hello Gunwant ke launde")
	}
	var welcome=greet
	welcome()
}