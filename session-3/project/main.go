package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if we have enough arguments
	if len(os.Args) < 5 {
		PrintUsage()
		os.Exit(1)
	}

	// Parse command-line arguments
	operation := os.Args[1]
	cipherMethod := os.Args[2]
	inputFile := os.Args[3]
	outputFile := os.Args[4]

	// Validate operation
	if operation != "encrypt" && operation != "decrypt" {
		fmt.Println("Error: Operation must be 'encrypt' or 'decrypt'")
		PrintUsage()
		os.Exit(1)
	}

	// Validate cipher method
	if cipherMethod != "rot13" && cipherMethod != "key" {
		fmt.Println("Error: Cipher method must be 'rot13' or 'key'")
		PrintUsage()
		os.Exit(1)
	}

	// Check if key is provided for key cipher
	var key string
	if cipherMethod == "key" {
		if len(os.Args) < 6 {
			fmt.Println("Error: Key is required for key cipher")
			PrintUsage()
			os.Exit(1)
		}
		key = os.Args[5]
	}

	// Read input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Process the content
	var result string
	if operation == "encrypt" {
		if cipherMethod == "rot13" {
			result = EncryptRot13(string(content))
		} else { // key cipher
			result = EncryptKey(string(content), key)
		}
	} else { // decrypt
		if cipherMethod == "rot13" {
			result = EncryptRot13(string(content)) // ROT13 is its own inverse
		} else { // key cipher
			result = DecryptKey(string(content), key)
		}
	}

	// Write to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully %sed the file using %s cipher.\n", operation, cipherMethod)
	fmt.Printf("Output written to: %s\n", outputFile)
}
