package main

import "fmt"

func displayNumber() func() int {
	number:=0

	return func() int {
		number=number+1
		return number
	}
}

func main() {
	num1:=displayNumber()
	fmt.Println(displayNumber()) //this returns the address of the anonymous function
	fmt.Println(displayNumber()())//this return the result of function,but here the value is not stored. the number is initialized every time
	fmt.Println(displayNumber()())
	fmt.Println(num1())
	fmt.Println(num1())
}