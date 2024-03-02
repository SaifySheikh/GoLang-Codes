package main

import "fmt"

func main() {
    var t int =3

	switch t{
		case 1,2,3: 
					fmt.Println("quarter 1")
					fallthrough
		case 4,5,6: 
					fmt.Println("quarter 2")
					fallthrough
					
		default: fmt.Print("Invalid")
		}
}
