package main

import "fmt"

func Findquare(n int)int{
	return n*n
}

func calsum(n1,n2 int)int{
	return n1+n2
}

func main() {
	// var sum=func(n1,n2 int)int{
	// 	sum:=n1+n2
	// 	return sum
	// }
	// result:=Findquare(sum(2,3))
	result:=Findquare(calsum(2,3))
	fmt.Println("result is ", result)
	
	//data type of sum is func
}