package main

import "fmt"

// Function that takes two strings and returns their combination
func combineStrings(str1, str2 string) string {
	return str1 + str2
}

// Function with multiple return values
func divideAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	// Example of using a function with string parameters
	result := combineStrings("Hello, ", "Functions!")
	fmt.Println(result)

	// Example of using a function with multiple return values
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Printf("10 divided by 3: Quotient = %d, Remainder = %d\n", quotient, remainder)
}

/* 
Tasks:
1. Create a function called 'multiply' that takes two integers and returns their product
2. Create a function called 'greet' that takes a name as a parameter and returns a greeting string
3. Bonus: Create a function with multiple return values

Concepts covered:
- Function declaration syntax: func functionName(parameters) returnType
- Multiple parameters of the same type: func example(x, y int)
- Return values and their types
- Multiple return values using parentheses: (int, int)
- Function calls and assignment of return values
*/
