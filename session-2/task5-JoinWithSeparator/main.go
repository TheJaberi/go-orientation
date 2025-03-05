package main

import (
	"fmt"
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
	words := []string{"Go", "is", "fun"}
	separator := "-"
	joined := JoinWithSeparator(words, separator)
	fmt.Printf("Question 5: Join Words with a Separator\n")
	fmt.Printf("Input: %v, Separator: %q\n", words, separator)
	fmt.Printf("Output: %q\n", joined)
}
// Question 5: Join Words with a Separator
// Input: ["Go", "is", "fun"], Separator: "-"
// Output: "Go-is-fun"
