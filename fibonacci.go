package main

import (
    "fmt"
    "strconv"
)


func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
    for i := 0; i <= 10; i++ {
        fmt.Print(strconv.Itoa(Fibonacci(i)) + " ")
    }
    fmt.Println("")
}
