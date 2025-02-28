package main

import (
	"fmt"
	"strings"
)

// CaseInsensitiveCompare compares two strings ignoring case.
func CaseInsensitiveCompare(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

// func toLower(r rune) rune {
// 	if r >= 'A' && r <= 'Z' {
// 		return r + ('a' - 'A')
// 	}
// 	return r
// }

// func CaseInsensitiveCompare(s1, s2 string) bool {
// 	runes1 := []rune(s1)
// 	runes2 := []rune(s2)

// 	if len(runes1) != len(runes2) {
// 		return false
// 	}

// 	for i, r := range runes1 {
// 		if toLower(r) != toLower(runes2[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }


func main() {
	s1 := "Hello"
	s2 := "hello"
	result := CaseInsensitiveCompare(s1, s2)
	fmt.Printf("Question 2: Compare Two Strings (Case-Insensitive)\n")
	fmt.Printf("Input: %q and %q\n", s1, s2)
	fmt.Printf("Output: %v\n", result)
}
//Compare Two Strings (Case-Insensitive) Without Using strings Package
// Input: "Hello" and "hello"
// Output: true
