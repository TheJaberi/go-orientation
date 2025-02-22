package main

import "fmt"

// Convert a string of digits to an integer
func stringToNumber(str string) int {
    result := 0
    
    // Loop through each character
    for i := 0; i < len(str); i++ {
        // Convert char to int by subtracting '0'
        // '0' is 48 in ASCII, so '1' - '0' = 1
        digit := int(str[i] - '0')
        
        // Multiply previous result by 10 and add new digit
        result = result*10 + digit
    }
    
    return result
}

func main() {
    // Example conversions
    number1 := "123"
    result1 := stringToNumber(number1)
    fmt.Printf("Converting %s to number: %d\n", number1, result1)
    
    number2 := "456"
    result2 := stringToNumber(number2)
    fmt.Printf("Converting %s to number: %d\n", number2, result2)
}

/* 
Task:
1. Convert a string of digits into a number
   Example: "123" → 123

Concepts covered:
- String to integer conversion
- ASCII values of characters
- Converting characters to numbers
- Building numbers digit by digit
*/
