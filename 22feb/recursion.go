package main

import "fmt"

func countdown(n int) {

	if(n==0){
		return
	}

	fmt.Println(n)

	countdown(n-1)

}

// func main() {
// 	countdown(5)
// }