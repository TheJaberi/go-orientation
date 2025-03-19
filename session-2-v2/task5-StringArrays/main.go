package main

/*
Task 5: String Arrays
-------------
Description:
    Join strings with separator and split strings into arrays.

Expected Input:
    Array of strings and a separator

Expected Output:
    Joined string and split results

Test Cases:
    1. ["a", "b", "c"], "," -> "a,b,c"
    2. "x-y-z", "-" -> ["x", "y", "z"]

$ go run main.go "," "a" "b" "c"
*/

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <separator> <string1> [string2...]")
        fmt.Println("Example: go run main.go \",\" \"a\" \"b\" \"c\"")
        os.Exit(1)
    }

    separator := os.Args[1]
    elements := os.Args[2:]

    // Join strings
    joined := strings.Join(elements, separator)
    fmt.Printf("Joined with '%s': %s\n", separator, joined)

    // Split string
    split := strings.Split(joined, separator)
    fmt.Printf("Split result: %v\n", split)
}
