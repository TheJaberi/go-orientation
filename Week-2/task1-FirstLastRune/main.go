package main

import (
	"fmt"
)

// FirstLastRune returns the first and last rune of the given string.
// If the string is empty, it returns 0 for both runes.
func FirstLastRune(s string) (rune, rune) {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0, 0
	}
	return runes[0], runes[len(runes)-1]
}

func main() {
	input := "Hello"
	first, last := FirstLastRune(input)
	fmt.Printf("Question 1: Extract First and Last Rune\n")
	fmt.Printf("Input: %q\n", input)
	fmt.Printf("Output: First rune: %c, Last rune: %c\n", first, last)
}
// Extract First and Last Rune
// Input: "Hello"
// Output: First rune: H, Last rune: o