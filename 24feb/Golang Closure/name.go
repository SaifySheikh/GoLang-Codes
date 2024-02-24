package main

import "fmt"

func printname() func() string{
	name:="Saify"
	// This anonymous func can't be used in main
	return func() string{ 
		name="Hi "+name
		return name
	}
}


func main(){
	message:=printname()

	fmt.Println(message())
	fmt.Println(message())
	fmt.Println(message())

}