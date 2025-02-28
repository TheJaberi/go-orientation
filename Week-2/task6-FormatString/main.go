package main

import (
	"fmt"
)


// FormatString transforms the input string such that:
// - The first half of the string is converted to uppercase.
// - The second half is converted to lowercase.
// - An underscore is inserted in the middle.
// If the length is odd, the underscore is placed after the middle character.
func FormatString(s string) string {
	runes := []rune(s)
	n := len(runes)

	var mid int
	if n%2 == 0 {
		mid = n / 2
	} else {
		mid = n/2 + 1
	}

	var result []rune

	for i := 0; i < mid; i++ {
		result = append(result, toUpper(runes[i]))
	}

	result = append(result, '_')

	for i := mid; i < n; i++ {
		result = append(result, toLower(runes[i]))
	}

	return string(result)
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func main() {
	testCases := []string{
		"golang",    // even length; expected: "GOL_ang"
		"helloWorld",// even length; expected: "HELLO_world"
		"test",      // even length; expected: "TE_st"
		"abcde",     // odd length; expected: "ABC_de"
	}

	fmt.Println("Question 6: Format String with Uppercase, Lowercase, and Underscore")
	for _, input := range testCases {
		output := FormatString(input)
		fmt.Printf("Input: %q -> Output: %q\n", input, output)
	}
}
