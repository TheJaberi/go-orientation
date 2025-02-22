package main

import "fmt"

// Basic calculator function
func calculate(a, b int) {
    // Addition
    fmt.Printf("%d + %d = %d\n", a, b, a+b)
    
    // Subtraction
    fmt.Printf("%d - %d = %d\n", a, b, a-b)
    
    // Multiplication
    fmt.Printf("%d * %d = %d\n", a, b, a*b)
    
    // Division (with check for divide by zero)
    if b != 0 {
        fmt.Printf("%d / %d = %d\n", a, b, a/b)
        fmt.Printf("%d %% %d = %d\n", a, b, a%b)  // Modulus (remainder)
    } else {
        fmt.Println("Cannot divide by zero")
    }
}

func main() {
    // Example with two numbers
    fmt.Println("Calculator with 10 and 3:")
    calculate(10, 3)
    
    fmt.Println("\nCalculator with 15 and 5:")
    calculate(15, 5)
}

/* 
Task:
1. Use basic arithmetic operators to perform calculations
   - Addition (+)
   - Subtraction (-)
   - Multiplication (*)
   - Division (/)
   - Modulus (%)

Concepts covered:
- Basic arithmetic operators
- Integer operations
- Division and remainder
- Basic error checking (divide by zero)
*/
