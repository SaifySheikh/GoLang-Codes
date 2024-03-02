package main

import "fmt"

func main() {
    var time float64
    fmt.Println("Enter Time")
    fmt.Scanf("%f", &time)

    if time > 0 && time < 25 {
        if time <= 10 {
            fmt.Println("Good Morning")
        } else if time < 20 && time > 10 {
            fmt.Println("Good day")
        } else {
            fmt.Println("Good Afternoon")
        }
    }
}
