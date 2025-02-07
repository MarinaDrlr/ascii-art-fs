package main

import (
	"fmt"
	"os"
)

// GetInput retrieves the input text and chosen font
func GetInput() (string, string, error) {
	if len(os.Args) < 2 {
		return "", "", fmt.Errorf("no input provided")
	}

	input := os.Args[1] // First argument: the text to be converted
	font := "standard"  // Default font

	if len(os.Args) > 2 {
		font = os.Args[2] // Second argument (if provided): the font type
	}

	return input, font, nil
}
