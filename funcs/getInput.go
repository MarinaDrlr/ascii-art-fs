package funcs

import (
	"fmt"
	"os"
)

// GetInput retrieves the user's input text and the chosen font (or defaults to "standard").
func GetInput() (string, string, error) {
	if len(os.Args) < 2 {
		return "", "", fmt.Errorf("no input provided")
	}

	input := os.Args[1] // User's input text
	font := "standard"  // Default font if none is provided

	if len(os.Args) > 2 {
		font = os.Args[2] // User-specified font (if provided)
	}

	return input, font, nil // Return the user input, chosen font, and no error
}
