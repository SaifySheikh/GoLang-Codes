package main
import "fmt"

func main(){
	var year int = 0
	fmt.Printf("Enter number:\n")
	fmt.Scanf("%d",&year)

	if(year == 0){
		fmt.Println("Please enter a valid number")
	}else{
		if((year%4==0 && year%100!=0) || (year%4==0 && year%100==0 && year%400==0)){
			fmt.Printf("Entered number is a leap year")
		}else{
			fmt.Printf("Not a leap year")
		}
	}
}