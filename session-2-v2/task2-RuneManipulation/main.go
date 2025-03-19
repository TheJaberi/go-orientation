package main

/*
Task 2: Rune Manipulation
-------------
Description:
    Extract and convert runes from a string to their numeric values.

Expected Input:
    A string containing any valid UTF-8 characters

Expected Output:
    First and last runes with their numeric values

Test Cases:
    1. "Hello" -> First: H (72), Last: o (111)
    2. "世界" -> First: 世 (19990), Last: 界 (30028)

$ go run main.go "Hello世界"
*/

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <string>")
        os.Exit(1)
    }

    input := os.Args[1]
    if len(input) == 0 {
        fmt.Println("Error: empty string provided")
        os.Exit(1)
    }

    runes := []rune(input)
    first, last := runes[0], runes[len(runes)-1]

    fmt.Printf("First rune: %c (Unicode: %d)\n", first, first)
    fmt.Printf("Last rune: %c (Unicode: %d)\n", last, last)
}
