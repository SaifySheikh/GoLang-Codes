package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	fmt.Println(today)

	fmt.Println(today.Format)
	fmt.Println(today.Day())
	fmt.Println(today.Year())

	switch today.Day(){
	case 1: fmt.Println("today is 1st")
	case 2: fmt.Println("today is 2nd")
	default: fmt.Println("Default")
	}
	
}
