package main
import "fmt"

func calSum(n int) int{
	if(n==0){
		return 0 
	}else{
		return n+calSum(n-1)
	}
}

func main() {
	var num int
	fmt.Println("Enter a number : ")
	fmt.Scan(&num)

	var result=calSum(num)
	fmt.Println(result)

}	