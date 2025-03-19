/*
Question 4: Convert Number to Base Representation
--------------------------------------------------
Write a Go program that converts an integer to its string representation in a specified base. 
You should handle negative numbers and use the provided character set for representing the digits. 
For example, for base 16, the digits would be "0123456789ABCDEF".

Example:
Input: Number = 255, Base = "0123456789ABCDEF"
Output: "FF"

The program should return the number as a string representation in the specified base.
*/

package main

import (
	"fmt"
)

// ConvertToBase converts an integer to its string representation in a specified base.
func ConvertToBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	negative := n < 0
	if negative {
		n = -n
	}

	b := len(base)
	var result []byte

	// Convert number to the specified base
	for n > 0 {
		remainder := n % b
		result = append([]byte{base[remainder]}, result...)
		n /= b
	}

	// Add negative sign if needed
	if negative {
		result = append([]byte{'-'}, result...)
	}

	return string(result)
}

func main() {
	number := 255
	base := "0123456789ABCDEF"
	converted := ConvertToBase(number, base)

	fmt.Println("Question 4: Convert Number to Base Representation")
	fmt.Printf("Input: Number = %d, Base = %q\n", number, base)
	fmt.Printf("Output: %q\n", converted)
}
