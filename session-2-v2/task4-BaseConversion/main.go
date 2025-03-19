package main

/*
Task 4: Base Conversion
-------------
Description:
    Convert numbers between different bases (binary, decimal, hex).

Expected Input:
    A number string and target base (2, 10, or 16)

Expected Output:
    The number converted to the specified base

Test Cases:
    1. "0xFF" 2 -> "0b11111111"
    2. "42" 16 -> "0x2A"

$ go run main.go "0xFF" 2
*/

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <number> <target_base>")
        fmt.Println("Example: go run main.go \"0xFF\" 2")
        os.Exit(1)
    }

    // Parse input number in any base
    input := strings.ToLower(os.Args[1])
    var num int64
    var err error

    switch {
    case strings.HasPrefix(input, "0b"):
        num, err = strconv.ParseInt(input[2:], 2, 64)
    case strings.HasPrefix(input, "0x"):
        num, err = strconv.ParseInt(input[2:], 16, 64)
    default:
        num, err = strconv.ParseInt(input, 10, 64)
    }

    if err != nil {
        fmt.Printf("Error parsing number: %v\n", err)
        os.Exit(1)
    }

    // Convert to target base
    base, err := strconv.Atoi(os.Args[2])
    if err != nil || (base != 2 && base != 10 && base != 16) {
        fmt.Println("Base must be 2, 10, or 16")
        os.Exit(1)
    }

    // Format result
    var result string
    switch base {
    case 2:
        result = fmt.Sprintf("0b%b", num)
    case 16:
        result = fmt.Sprintf("0x%X", num)
    default:
        result = fmt.Sprintf("%d", num)
    }

    fmt.Printf("Result: %s\n", result)
}
