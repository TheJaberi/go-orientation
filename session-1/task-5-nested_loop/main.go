package main

import "fmt"

// Function to print a square pattern
func printSquare(size int) {
    // Outer loop controls rows
    for i := 0; i < size; i++ {
        // Inner loop controls columns
        for j := 0; j < size; j++ {
            fmt.Print("* ")
        }
        // After each row, print a new line
        fmt.Println()
    }
}

func main() {
    fmt.Println("Printing a 3x3 square:")
    printSquare(3)
    
    fmt.Println("\nPrinting a 5x5 square:")
    printSquare(5)
}

/* 
Task:
1. Create a program that prints a rectangle pattern using nested loops
   Example output for size 3:
   * * *
   * * *
   * * *

Concepts covered:
- Nested loops - one loop inside another
- Outer loop controls rows
- Inner loop controls columns
- Using loops to create patterns
*/
