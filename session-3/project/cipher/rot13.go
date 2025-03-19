package cipher

import "strings"

// EncryptRot13 applies ROT13 encryption/decryption to the input string
// ROT13 is its own inverse, so the same function works for both encryption and decryption
func EncryptRot13(input string) string {
	var result strings.Builder
	for _, char := range input {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			// Determine the base (lowercase or uppercase)
			base := 'a'
			if char >= 'A' && char <= 'Z' {
				base = 'A'
			}

			// Apply ROT13 transformation
			rotated := (char-base+13)%26 + base
			result.WriteRune(rotated)
		} else {
			// Non-alphabetic characters remain unchanged
			result.WriteRune(char)
		}
	}
	return result.String()
}
