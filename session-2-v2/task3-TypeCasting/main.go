package main

/*
Task 3: Type Casting
-------------
Description:
    Demonstrate type conversions between different data types.

Expected Input:
    Integer, float, and string numbers

Expected Output:
    Values converted between different types

Test Cases:
    1. "42" -> int: 42, float64: 42.0
    2. 3.14 -> int: 3, string: "3.14"

$ go run main.go 42 3.14 "123"
*/

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run main.go <integer> <float> <string>")
        os.Exit(1)
    }

    // Integer conversions
    intVal, err := strconv.ParseInt(os.Args[1], 10, 64)
    if err != nil {
        fmt.Printf("Error converting to int: %v\n", err)
        os.Exit(1)
    }

    // Float conversions
    floatVal, err := strconv.ParseFloat(os.Args[2], 64)
    if err != nil {
        fmt.Printf("Error converting to float: %v\n", err)
        os.Exit(1)
    }

    // Display conversions
    fmt.Printf("\nInteger conversions:\n")
    fmt.Printf("%d -> float64: %f\n", intVal, float64(intVal))
    fmt.Printf("%d -> string: %s\n", intVal, strconv.FormatInt(intVal, 10))

    fmt.Printf("\nFloat conversions:\n")
    fmt.Printf("%f -> int64: %d\n", floatVal, int64(floatVal))
    fmt.Printf("%f -> string: %s\n", floatVal, strconv.FormatFloat(floatVal, 'f', 2, 64))
}
