package funcs

import (
	"fmt"
	"os"
)

// GetInput retrieves the user's input string and banner font.
// It requires exactly 2 arguments: the input text and the font name.
// If the number of arguments is not exactly 2, it returns a usage error.
func GetInput() (string, string, error) {
	args := os.Args[1:] // Skip the program name

	if len(args) != 2 {
		// Invalid number of arguments – return usage message
		return "", "", fmt.Errorf("Usage: go run . [STRING] [BANNER]\nEX: go run . \"something\" standard")
	}

	input := args[0] // Required: the text to convert to ASCII art
	font := args[1]  // Required: the font name (e.g., standard, shadow, thinkertoy)

	return input, font, nil
}
