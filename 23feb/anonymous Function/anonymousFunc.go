package main

import "fmt"

func main() {
	var greet=func(){
		fmt.Println("hello")
	}
	greet()
}