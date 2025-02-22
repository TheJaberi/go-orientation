package main

import "fmt"

// Function that demonstrates pass by value
func modifyValue(x int) {
    x = 100 // This change won't affect the original value
    fmt.Println("Inside function:", x)
}

// Function that demonstrates pass by reference
func modifyReference(x *int) {
    *x = 100 // This change will affect the original value
    fmt.Println("Inside function:", *x)
}

func main() {
    // Pass by value example
    fmt.Println("Pass by value example:")
    number := 42
    fmt.Println("Before:", number)
    modifyValue(number)
    fmt.Println("After:", number)
    
    // Pass by reference example
    fmt.Println("\nPass by reference example:")
    number = 42
    fmt.Println("Before:", number)
    modifyReference(&number)
    fmt.Println("After:", number)
}

/* 
Task:
1. Create two functions that show the difference between:
   - Passing a value (makes a copy)
   - Passing a reference (can modify original)

Concepts covered:
- Pass by value: function gets a copy
- Pass by reference: function can modify original
- When to use each method
- Understanding parameter passing
*/
