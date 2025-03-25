package funcs

import (
	"fmt"
	"os"
)

// // GetInput retrieves the user's input text and the chosen font (or defaults to "standard").
// func GetInput() (string, string, error) {
// 	if len(os.Args) < 2 {
// 		return "", "", fmt.Errorf("no input provided")
// 	}

// 	input := os.Args[1] // User's input text
// 	font := "standard"  // Default font if none is provided

// 	if len(os.Args) > 2 {
// 		font = os.Args[2] // User-specified font (if provided)
// 	}

// 	return input, font, nil // Return the user input, chosen font, and no error
// }

// GetInput retrieves the user's input text and the chosen font (or defaults to "standard").
// It expects 1 or 2 arguments (excluding flags). If the argument count is invalid, it returns an error.
func GetInput() (string, string, error) {
	args := os.Args[1:] // Skip the program name

	// Validate the number of non-flag arguments
	if len(args) == 0 || len(args) > 2 {
		return "", "", fmt.Errorf("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}

	input := args[0]   // Required: the text to convert to ASCII art
	font := "standard" // Default font

	if len(args) == 2 {
		font = args[1] // Optional: user-specified font
	}

	return input, font, nil
}
