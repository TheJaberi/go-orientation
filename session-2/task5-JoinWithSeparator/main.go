/*
Question 5: Join Words with a Separator
----------------------------------------
Write a Go program that accepts a separator string and a list of words as command-line arguments.
The program should join the list of words using the separator and display the result.

Usage:
go run main.go <separator> <word1> <word2> ... <wordN>

Example:
Input: Words = ["apple", "banana", "cherry"], Separator = ", "
Output: "apple, banana, cherry"
*/

package main

import (
	"fmt"
	"os"
)

// JoinWithSeparator joins a slice of strings using the provided separator.
func JoinWithSeparator(words []string, separator string) string {
	if len(words) == 0 {
		return ""
	}

	result := words[0]
	for i := 1; i < len(words); i++ {
		result += separator + words[i]
	}
	return result
}

func main() {
	// Ensure there are enough command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <separator> <word1> <word2> ... <wordN>")
		return
	}

	// Extract separator and words from arguments
	separator := os.Args[1]
	words := os.Args[2:]

	// Join words using the provided separator
	result := JoinWithSeparator(words, separator)

	// Display the question and results clearly
	fmt.Println("Question 5: Join Words with a Separator")
	fmt.Printf("Input: Words = %v, Separator = %q\n", words, separator)
	fmt.Printf("Output: %q\n", result)
}
