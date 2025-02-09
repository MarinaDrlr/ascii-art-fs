package main

import (
	"ascii-art/funcs"
	"fmt"
	"os"
)

func main() {
	// Get user input and font choice
	input, font, err := funcs.GetInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Normalize the input to match the banner map keys
	normalizedInput := funcs.NormalizeInput(input)

	// Check if the normalized input is empty
	if normalizedInput == "" {
		// fmt.Println("Error: Invalid or empty input provided. Please try again.") // Debugging message (commented out)
		return // Simply exit without printing anything
	}

	// Load the banner file dynamically
	banner, err := funcs.LoadBanner(font)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Generate ASCII art
	asciiArt, err := funcs.GenerateASCIIArt(normalizedInput, banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// If ASCII art is empty but input was just newlines, print a blank line
	if len(asciiArt) == 0 {
		fmt.Println() // Prints a blank line
		return
	}

	// Display the result
	funcs.DisplayASCIIArt(asciiArt)

	// fmt.Println("ASCII art generated successfully.")
}
