/*
Question: Format Strings from File
----------------------------------
Write a Go program that reads words from an input file, processes each word according to the following format, and writes the formatted results to an output file:
1. The first half of each word should be converted to uppercase.
2. The second half should be converted to lowercase.
3. An underscore should be inserted in the middle. If the length of the word is odd, the underscore is placed after the middle character.

For example:
Input: "hello", "world"
Output: "HE_llo", "WO_rld"

The program should handle multiple words per line and write the results to a file.

Usage:
go run main.go
Input file: input.txt
Output file: result.txt
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// Convert the first half to uppercase
	for i := 0; i < mid; i++ {
		result = append(result, toUpper(runes[i]))
	}

	// Add the underscore
	result = append(result, '_')

	// Convert the second half to lowercase
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
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create or open the output file for writing
	resultFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Error creating result file:", err)
		return
	}
	defer resultFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the file
	for scanner.Scan() {
		// For each line, split it into words
		words := strings.Fields(scanner.Text())

		// Process each word and write the formatted result to the file
		for _, word := range words {
			output := FormatString(word)
			_, err = fmt.Fprint(resultFile, output+" ") // Write only the transformed word
			if err != nil {
				fmt.Println("Error writing to result file:", err)
				return
			}
		}
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
