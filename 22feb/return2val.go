package main

import "fmt"

func calculateSum(n1 int,n2 int) (int,int) {
	sum := n1+n2
	difference := n1-n2
	fmt.Println("sum is ", sum)
	fmt.Println("difference is ", difference)

	return sum,difference

}

// func main(){

// 	fmt.Println("Enter the value of the numbers")

// 	var num1 int
// 	var num2 int
// 	var result1 int
// 	var result2 int

// 	fmt.Scan(&num1, &num2)
// 	result1,result2=CalculateSum(num1,num2)

// 	fmt.Println(result1,result2)

// }
