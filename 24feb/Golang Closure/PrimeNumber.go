package main

import "fmt"

func calculate() func() int{
	num:=1
	// This anonymous func can't be used in main
	return func() int{ 
		num=num+2
		return num
	}
}


func main(){
	odd:=calculate() //it's like creating an object of alculate
	fmt.Println(odd()) //value of num is stored for odd
	fmt.Println(odd())
	fmt.Println(odd())

	odd1:=calculate() //it's like creating an object of calculate
	fmt.Println(odd1()) //value of num is stored for odd1

}