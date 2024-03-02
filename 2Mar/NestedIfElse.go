package main

import "fmt"

func main(){
	a:=10
	b:=20

	if(a!=b){
		if(a<b){
			fmt.Println("b is too big")
		}else{
			fmt.Println("a is too small")
		}

	}else{
		fmt.Println("a is equal to b")
	}
}