package main

import "fmt"

func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}

func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {

    fmt.Println("Factorial of 5:", Factorial(5))
    
    fmt.Println("Fibonacci sequence:")
    for i := 0; i < 10; i++ {
        fmt.Print(Fibonacci(i), " ")
    }
    fmt.Println()
}