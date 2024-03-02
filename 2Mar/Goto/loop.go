package main
import "fmt"


func main(){
	i:=0

	loop:
		fmt.Println(i)

		i++
		if(i<5){
			goto loop

		}

		fmt.Println("Loops ends here")
}