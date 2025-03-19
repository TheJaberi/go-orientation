/*
Question 3: Count Alphabetic Characters
-----------------------------------------
Write a Go program that counts the number of alphabetic characters (A-Z, a-z) in a given string. 
Ignore any non-alphabetic characters, such as numbers or punctuation.

Example:
Input: "Hello123!"
Output: 5

The program should return the number of alphabetic characters in the input string.
*/

package main

import "fmt"

// CountAlphabetic counts the number of alphabetic characters (A-Z, a-z) in a string.
func CountAlphabetic(s string) int {
	count := 0
	for _, ch := range s {
		if ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	input := "Hello123!"
	count := CountAlphabetic(input)

	fmt.Printf("Question 3: Count Alphabetic Characters\n")
	fmt.Printf("Input: %q\n", input)
	fmt.Printf("Output: %d alphabetic characters\n", count)
}
