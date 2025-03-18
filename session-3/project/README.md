# Session 3 Project: Secret Keeper

## Overview

Secret Keeper is a command-line tool that can encrypt and decrypt text files using two different cipher methods. This project demonstrates file I/O operations, command-line argument processing, and string manipulation in Go.

## Task Description

Your task is to implement a program that can encrypt and decrypt text files using two different cipher methods:

1. **ROT13** - A simple substitution cipher that replaces each letter with the letter 13 positions after it in the alphabet
2. **Key Cipher** - A more secure cipher where each character is shifted based on the corresponding character in a secret key

The program should:

1. Take command-line arguments to specify:
   - Operation mode (encrypt/decrypt)
   - Cipher method (rot13/key)
   - Input file path
   - Output file path
   - Secret key (required for key cipher, optional for ROT13)

2. Read content from the input file
3. Apply the selected encryption/decryption method
4. Write the processed content to the output file

## Project Structure

The project is organized into multiple files to demonstrate code modularity and separation of concerns:

```
session-3/
└── project/
    ├── main.go       # Main program flow
    ├── utils.go      # Utility functions
    ├── rot13.go      # ROT13 cipher implementation
    ├── keycipher.go  # Key-based cipher implementation
    ├── message.txt   # Example input file
    └── README.md     # Instructions
```

This structure helps keep the code organized and makes it easier to maintain. Each file has a specific purpose:

- **main.go**: Contains the main program flow, argument parsing, and file I/O
- **utils.go**: Contains utility functions like PrintUsage
- **rot13.go**: Contains the ROT13 cipher implementation
- **keycipher.go**: Contains the key-based cipher implementation

## Requirements

- Implement proper error handling for file operations and invalid arguments
- Preserve case (uppercase/lowercase) during encryption/decryption
- Non-alphabetic characters should remain unchanged
- For the key cipher, if the key is shorter than the input, it should repeat
- Organize code into multiple files for better modularity

## Usage

```
# Using ROT13 cipher
go run *.go encrypt rot13 [input_file] [output_file]
go run *.go decrypt rot13 [input_file] [output_file]

# Using key-based cipher
go run *.go encrypt key [input_file] [output_file] [secret_key]
go run *.go decrypt key [input_file] [output_file] [secret_key]
```

## Examples

```
# Encrypt a message using ROT13
go run *.go encrypt rot13 message.txt encrypted_rot13.txt

# Decrypt a ROT13-encrypted message
go run *.go decrypt rot13 encrypted_rot13.txt decrypted_rot13.txt

# Encrypt a message using key cipher
go run *.go encrypt key message.txt encrypted_key.txt mysecretkey

# Decrypt a key-encrypted message
go run *.go decrypt key encrypted_key.txt decrypted_key.txt mysecretkey
```

## Encryption Methods Explained

### ROT13 Cipher

ROT13 (rotate by 13 places) is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet. For example:

- 'A' becomes 'N'
- 'B' becomes 'O'
- 'Z' becomes 'M'

ROT13 is its own inverse, meaning applying it twice returns the original text. It's a weak encryption method but useful for learning purposes.

### Key Cipher

The key cipher is a variation of the Caesar cipher where the shift value for each character is determined by the corresponding character in the secret key. This makes it more secure than a simple ROT13 cipher.

For example, if the key is "KEY" and the message is "HELLO":
- 'H' is shifted by the value of 'K' (10)
- 'E' is shifted by the value of 'E' (4)
- 'L' is shifted by the value of 'Y' (24)
- 'L' is shifted by the value of 'K' (10) (key repeats)
- 'O' is shifted by the value of 'E' (4) (key repeats)

## Learning Objectives

- Working with command-line arguments in Go
- Reading from and writing to files
- String and rune manipulation
- Implementing encryption algorithms
- Error handling
- Function organization
- Code modularity and package structure

