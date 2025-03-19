package main

/*
Task 3: Command Line Arguments
-------------
Description:
    Write a program that demonstrates advanced command line argument processing
    including flags, positional arguments, and error handling.

Objectives:
    1. Learn to use the flag package for argument parsing
    2. Handle positional and optional arguments
    3. Implement proper error checking
    4. Process different argument types
    5. Provide helpful usage information

Expected Input:
    Command line flags and arguments:
    -name: string flag for name
    -age: integer flag for age
    -verbose: boolean flag for verbose output
    [arguments...]: Additional positional arguments

Expected Output:
    Processed command line arguments showing:
    - Parsed flag values
    - Positional arguments
    - Verbose information if enabled

Example Usage:
    $ go run main.go -name "John" -age 25 -verbose arg1 arg2
    Name: John
    Age: 25
    Verbose mode: enabled
    Additional arguments: [arg1 arg2]

Test Cases:
    1. Basic: go run main.go -name "John" -age 25
    2. Verbose: go run main.go -name "John" -age 25 -verbose
    3. With args: go run main.go -name "John" -age 25 file1.txt file2.txt
    4. Invalid age: go run main.go -name "John" -age abc
*/

import (
    "flag"
    "fmt"
    "os"
)

type Config struct {
    name    string
    age     int
    verbose bool
}

func parseFlags() (Config, []string) {
    cfg := Config{}

    // Define flags
    flag.StringVar(&cfg.name, "name", "", "Your name")
    flag.IntVar(&cfg.age, "age", 0, "Your age")
    flag.BoolVar(&cfg.verbose, "verbose", false, "Enable verbose output")

    // Custom usage message
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [flags] [arguments...]\n\nFlags:\n", os.Args[0])
        flag.PrintDefaults()
    }

    // Parse flags
    flag.Parse()

    // Get remaining arguments
    args := flag.Args()

    // Validate required flags
    if cfg.name == "" {
        fmt.Println("Error: -name flag is required")
        flag.Usage()
        os.Exit(1)
    }

    if cfg.age <= 0 {
        fmt.Println("Error: -age flag is required and must be positive")
        flag.Usage()
        os.Exit(1)
    }

    return cfg, args
}

func main() {
    // Parse command line flags and arguments
    config, remainingArgs := parseFlags()

    // Print configuration
    fmt.Printf("\nConfiguration:\n")
    fmt.Printf("Name: %s\n", config.name)
    fmt.Printf("Age: %d\n", config.age)
    
    if config.verbose {
        fmt.Println("\nVerbose output:")
        fmt.Printf("Program name: %s\n", os.Args[0])
        fmt.Printf("Number of arguments: %d\n", len(remainingArgs))
    }

    // Print remaining arguments if any
    if len(remainingArgs) > 0 {
        fmt.Printf("\nAdditional arguments: %v\n", remainingArgs)
    }
}
