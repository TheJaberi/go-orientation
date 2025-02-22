# Getting Started with Go

Welcome to Go programming! This guide will help you understand the basics before starting the tasks.

## What is Go?

Go is a programming language that's:
- Easy to learn
- Good at doing multiple things at once
- Has built-in tools that help you write better code

## Running Go Programs

1. Every Go file ends with `.go`
2. To run a program:
   ```bash
   go run filename.go
   ```

## Basic Program Structure

Every Go program looks like this:
```go
package main  // Tell Go this is a runnable program

import "fmt"  // Import packages we need

func main() {  // Program starts here
    // Your code goes here
}
```

## Variables

Making variables is easy:
```go
name := "John"           // Go figures out it's a string
age := 20               // Go figures out it's a number
message := "Hello!"     // := is for new variables

var score int = 100    // You can also be specific about the type
```

## Printing Things

Two main ways to print:
```go
fmt.Println("Hello!")   // Prints and adds a new line
fmt.Print("Hi ")        // Prints without a new line
```

Using `fmt.Printf` for formatted output:
```go
name := "John"
age := 20
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

## Packages

- Packages are like toolboxes with pre-made code
- `import "fmt"` gives you printing tools
- You'll use different packages for different tasks
- Your code must start with `package main` to be runnable

## Important Things to Remember

1. Every program needs:
   - `package main`
   - `func main()`
   - Proper `import` statements

2. Common mistakes to avoid:
   - Forgetting to put `package main` at the top
   - Forgetting curly braces `{}`
   - Not importing needed packages

3. Code style:
   - Go will automatically format your code
   - Use meaningful variable names
   - Add spaces to make code readable

## Basic Data Types

```go
text := "Hello"         // string (text)
number := 42           // int (whole number)
decimal := 3.14        // float64 (decimal number)
isTrue := true         // bool (true/false)
```

## Basic Operators

```go
sum := 5 + 3          // Addition
difference := 10 - 5   // Subtraction
product := 4 * 2      // Multiplication
quotient := 8 / 2     // Division
remainder := 7 % 3    // Remainder (modulus)
```

## Getting Help

- Comments in your code start with `//`
- Each task has comments explaining what to do
- Read error messages carefully - they tell you what's wrong

Now you're ready to start the tasks! Remember, programming is about trying things and learning from mistakes. Don't worry if you don't get everything right the first time!
