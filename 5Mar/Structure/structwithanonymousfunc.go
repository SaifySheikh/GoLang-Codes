package main

import "fmt"


type Rectangle func(int,int) int

type rectpara struct {
	length int
	breadth int
	color string
	rect Rectangle
}

func main(){
	rect1:= rectpara{
		length:10,
		breadth:10,
		color:"Red",
		rect:func(length, breadth int) int{
			return length*breadth
		},
	}

	fmt.Printf("%s\n", rect1.color)
	fmt.Printf("%d\n", rect1.length)
	fmt.Printf("%d\n", rect1.breadth)
	fmt.Printf("%d\n", rect1.rect(rect1.length, rect1.breadth))
	
}