package main

/*
Task 1: Struct Basics
-------------
Description:
    Create a Person struct with name and age fields and a method to display info.

Expected Input:
    Name string and age integer

Expected Output:
    Formatted person information

Test Cases:
    1. Person{"John", 25} -> "John is 25 years old"
    2. Person{"Alice", 30} -> "Alice is 30 years old"

$ go run main.go John 25
*/

import (
    "fmt"
    "os"
    "strconv"
)

type Person struct {
    name string
    age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s is %d years old", p.name, p.age)
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <name> <age>")
        os.Exit(1)
    }

    age, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error: age must be a number")
        os.Exit(1)
    }

    person := Person{
        name: os.Args[1],
        age:  age,
    }

    fmt.Println(person)
}
