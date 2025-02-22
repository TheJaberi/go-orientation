package main

import "fmt"

// Function that prints numbers from 1 to n
func printNumbers(n int) {
    for i := 1; i <= n; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()
}

// Function that demonstrates different loop patterns
func loopPatterns() {
    // Basic for loop with condition
    for i := 0; i < 3; i++ {
        fmt.Println("Count:", i)
    }

    // While-like loop
    sum := 0
    for sum < 10 {
        sum += 2
        fmt.Println("Sum is now:", sum)
    }

    // Infinite loop with break
    count := 0
    for {
        count++
        if count > 3 {
            break
        }
        fmt.Println("In infinite loop, count:", count)
    }
}

func main() {
    fmt.Println("Printing numbers 1 to 5:")
    printNumbers(5)
    
    fmt.Println("\nDemonstrating different loop patterns:")
    loopPatterns()
}

/* 
Tasks:
1. Write a program that prints numbers from 1 to 10 using a for loop
2. Create a loop that sums numbers from 1 to 5
3. Bonus: Create a loop that prints only even numbers from 0 to 10

Concepts covered:
- Basic for loop syntax: for initialization; condition; post
- While-like loops in Go: for condition { }
- Infinite loops: for { }
- Break statement to exit loops
- Loop counter variables
- Loop conditions and iterations
*/
