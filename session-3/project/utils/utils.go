package utils

import "fmt"

// PrintUsage prints the usage instructions
func PrintUsage() {
fmt.Println("Usage:")
fmt.Println("  ROT13 cipher: go run main.go [encrypt|decrypt] rot13 [input_file] [output_file]")
fmt.Println("  Key cipher:   go run main.go [encrypt|decrypt] key [input_file] [output_file] [secret_key]")
}
