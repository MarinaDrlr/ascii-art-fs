package main

import (
	"fmt"
	"os"
	"strings"
)

// GetInput retrieves input string and chosen font
func GetInput() (string, string) {
	if len(os.Args) < 2 {
		fmt.Println("Error: No input provided")
		os.Exit(1)
	}

	input := os.Args[1] // First argument: the text to be converted
	font := "standard"  // Default font

	if len(os.Args) == 3 {
		font = os.Args[2] // Second argument (if provided): the font type
	}

	return input, font
}

// NormalizeInput normalizes input (convert to uppercase or lowercase)
func NormalizeInput(input string) string {
	return strings.ToUpper(input) // Ensure case matches the banner file
}
