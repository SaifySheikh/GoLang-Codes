package main

import "fmt"

func main() {
    var t int =3

	// switch t{
	// case 1: fmt.Println("Jan")
	// case 2: fmt.Println("Feb")
	// case 3: fmt.Println("Mar")
	// default: fmt.Print("Invalid")
	// }

	switch t{
		case 1,2,3: fmt.Println("quarter 1")
		case 4,5,6: fmt.Println("quarter 2")
		default: fmt.Print("Invalid")
		}
}
