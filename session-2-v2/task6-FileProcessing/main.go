package main

/*
Task 6: File Processing
-------------
Description:
    Read text from a file and display its contents.

Expected Input:
    Path to an input file

Expected Output:
    File contents or error message

Test Cases:
    1. "input.txt" (exists) -> displays content
    2. "missing.txt" -> shows error

$ go run main.go input.txt
*/

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <input_file>")
        os.Exit(1)
    }

    content, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("File contents:\n%s", string(content))
}
