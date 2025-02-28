package main

import (
	"fmt"
)

// CountAlpha counts the number of alphabetic characters in a string.
func CountAlpha(s string) int {
	count := 0
	for _, ch := range s {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	input := "Hello123!"
	count := CountAlpha(input)
	fmt.Printf("Question 3: Count Alphabetic Characters\n")
	fmt.Printf("Input: %q\n", input)
	fmt.Printf("Output: %d alphabetic characters\n", count)
}
// Count Alphabetic Characters
// Input: "Hello123!"
// Output: 5 alphabetic characters
