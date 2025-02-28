package main

import (
	"fmt"
)

// ConvertToBase converts an integer to its string representation in a specified base.
func ConvertToBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	b := len(base)
	var result []byte
	for n > 0 {
		remainder := n % b
		result = append([]byte{base[remainder]}, result...)
		n /= b
	}
	if negative {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func main() {
	number := 255
	base := "0123456789ABCDEF"
	converted := ConvertToBase(number, base)
	fmt.Printf("Question 4: Convert Number to Base Representation\n")
	fmt.Printf("Input: %d, Base: %q\n", number, base)
	fmt.Printf("Output: %q\n", converted)
}
// Question 4: Convert Number to Base Representation
// Input: 255, Base: "0123456789ABCDEF"
// Output: "FF"
