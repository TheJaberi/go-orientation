/*
Question 1: Extract First and Last Rune
----------------------------------------
Write a Go program that extracts the first and last rune from a given string. If the string is empty, it should return 0 for both the first and last rune.

Question 2: Character Replacement
---------------------------------
Write a Go program that performs character replacement in a string based on the following mapping:
- 'A' -> '4', 'E' -> '3', 'I' -> '1', 'O' -> '0', 'S' -> '$'
- Lowercase versions of these characters should also be replaced accordingly.

Usage:
1. If command-line arguments are provided, process the input string and output the modified string.
2. If no arguments are provided, use a default string and display the results for both tasks.

Example for Character Replacement:
Input: "Hello"
Output: "H3ll0"

File Processing:
You can also process an input file with character replacements and write the modified result to an output file.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

// ReplaceCharacters replaces specific characters based on the given mapping.
func ReplaceCharacters(s string) string {
	replacements := map[rune]rune{
		'A': '4', 'E': '3', 'I': '1', 'O': '0', 'S': '$',
		'a': '4', 'e': '3', 'i': '1', 'o': '0', 's': '$',
	}

	runes := []rune(s)
	for i, r := range runes {
		if newVal, exists := replacements[r]; exists {
			runes[i] = newVal
		}
	}
	return string(runes)
}

// ProcessFile reads from an input file, applies character replacement, and writes to an output file.
func ProcessFile(inputFile, outputFile string) error {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	modified := ReplaceCharacters(string(data))
	return ioutil.WriteFile(outputFile, []byte(modified), 0644)
}

func main() {
	if len(os.Args) > 1 {
		// Process command-line input
		input := strings.Join(os.Args[1:], " ")
		fmt.Println("Modified String:", ReplaceCharacters(input))
	} else {
		// Default test case
		input := "Hello"
		first, last := FirstLastRune(input)
		fmt.Printf("Question 1: Extract First and Last Rune\n")
		fmt.Printf("Input: %q\n", input)
		fmt.Printf("Output: First rune: %c, Last rune: %c\n", first, last)

		// Test character replacement
		fmt.Println("Question 2: Character Replacement")
		fmt.Printf("Original: %s\n", input)
		fmt.Printf("Modified: %s\n", ReplaceCharacters(input))

		// File Processing Example (Uncomment to test with files)
		// err := ProcessFile("input.txt", "output.txt")
		// if err != nil {
		// 	fmt.Println("Error processing file:", err)
		// }
	}
}
