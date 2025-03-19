/*
Question 2: Compare Two Strings (Case-Insensitive)
---------------------------------------------------
Write a Go program that compares two strings in a case-insensitive manner. Do not use the `strings` package. 
You should convert the characters of both strings to lowercase before comparing them. 

Example:
Input: "Hello" and "hello"
Output: true

If the strings are the same (ignoring case), the program should return true; otherwise, it should return false.
*/

package main

import "fmt"

// toLower converts an uppercase letter to lowercase.
// If the rune is already lowercase or not a letter, it remains unchanged.
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

// CaseInsensitiveCompare compares two strings ignoring case, without using strings package.
func CaseInsensitiveCompare(s1, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	if len(runes1) != len(runes2) {
		return false
	}

	for i, r := range runes1 {
		if toLower(r) != toLower(runes2[i]) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "Hello"
	s2 := "hello"
	result := CaseInsensitiveCompare(s1, s2)

	fmt.Printf("Question 2: Compare Two Strings (Case-Insensitive)\n")
	fmt.Printf("Input: %q and %q\n", s1, s2)
	fmt.Printf("Output: %v\n", result)
}
