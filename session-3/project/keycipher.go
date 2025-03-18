package main

import "strings"

// EncryptKey applies key-based encryption to the input string
func EncryptKey(input string, key string) string {
	var result strings.Builder
	keyRunes := []rune(key)
	keyLength := len(keyRunes)
	
	for i, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Determine the base (lowercase or uppercase)
			base := 'a'
			if char >= 'A' && char <= 'Z' {
				base = 'A'
			}
			
			// Get the shift value from the key
			keyChar := keyRunes[i%keyLength]
			shift := int(keyChar % 26)
			
			// Apply the shift
			encrypted := (char-base+rune(shift))%26 + base
			result.WriteRune(encrypted)
		} else {
			// Non-alphabetic characters remain unchanged
			result.WriteRune(char)
		}
	}
	return result.String()
}

// DecryptKey applies key-based decryption to the input string
func DecryptKey(input string, key string) string {
	var result strings.Builder
	keyRunes := []rune(key)
	keyLength := len(keyRunes)
	
	for i, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Determine the base (lowercase or uppercase)
			base := 'a'
			if char >= 'A' && char <= 'Z' {
				base = 'A'
			}
			
			// Get the shift value from the key
			keyChar := keyRunes[i%keyLength]
			shift := int(keyChar % 26)
			
			// Apply the reverse shift (add 26 to ensure positive result)
			decrypted := (char-base+26-rune(shift))%26 + base
			result.WriteRune(decrypted)
		} else {
			// Non-alphabetic characters remain unchanged
			result.WriteRune(char)
		}
	}
	return result.String()
}
