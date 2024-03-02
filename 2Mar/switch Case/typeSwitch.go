package main
import "fmt"

func main(){
	var x interface{} = "RKNEC" //koi variable agr interface hai toh uska data type kabhi bhi change kr skte hai
	switch i := x.(type) {
		case string: fmt.Printf("string %T",i)
		case nil: fmt.Println("nil")
		case int64: fmt.Printf("int64")
		default: fmt.Println("unknown")
	}
}