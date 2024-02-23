package main

import "fmt"

func main() {
	var sum=func(n1,n2 int)int{
		sum:=n1+n2
		// fmt.Println("sum is ", sum)
		return sum

	}
	result:=sum(2,3)
	fmt.Println("result is ", result)
	fmt.Printf("Data type of sum is %T", sum)
	//data type of sum is func
}