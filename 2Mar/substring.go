package main

import "fmt"
import "strings"
	

func main(){
	var str string ="Saify"
	var substr string ="fy"

	if strings.Contains(str,substr) {
		fmt.Printf("%s Contains %s\n",str,substr)
	}else{
		fmt.Println("Doesn't contains")
	}


	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	fmt.Println(strings.Index(str, "y"))
}