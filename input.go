package main

import (
	"fmt"
	"os"
	"strings"
)

// Get input from command-line arguments
func getInput() string {
	if len(os.Args) < 2 {
		fmt.Println("Error: No input provided")
		os.Exit(1)
	}
	return os.Args[1]
}

// Normalize input (convert to uppercase or lowercase)
func normalizeInput(input string) string {
	return strings.ToUpper(input) // Adjust this depending on banner file requirements
}
