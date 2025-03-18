package main

import (
	"fmt"
	"os"
)

// JoinWithSeparator joins a slice of strings using the provided separator

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <separator> <word1> <word2> ... <wordN>")
		return
	}

	// The first argument is the separator
	separator := os.Args[1]

	// The remaining arguments are the words to join
	words := os.Args[2:]
	if len(words)==0{
		fmt.Println()
	}
	results:=words[0]
	for i:=1;i<len(words);i++{
		results+=separator+words[i]
	}
	fmt.Println(results)

	// Print the result
	fmt.Printf("Question 5: Join Words with a Separator\n")
	fmt.Printf("Input: %v, Separator: %q\n", words, separator)
	fmt.Printf("Output: %q\n", results)
}
