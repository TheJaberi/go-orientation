package main

import "fmt"

func main() {
    // Declare a variable
    x := 42
    
    // Create a pointer to x using &
    ptr := &x
    
    // Print original value
    fmt.Println("Original value:", x)
    
    // Print memory address
    fmt.Println("Memory address:", ptr)
    
    // Change value through pointer using *
    *ptr = 100
    
    // Print new value
    fmt.Println("New value:", x)
}

/* 
Task:
1. Create a variable and change its value using a pointer

Concepts covered:
- What is a pointer? (A variable that stores memory address)
- & operator: get memory address
- * operator: get value at memory address
- How to modify values using pointers
*/
