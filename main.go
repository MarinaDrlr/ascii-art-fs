package main

import "fmt"

func main() {
	// Get user input and font choice
	input, font := GetInput()

	// Normalize the input to match the banner map keys
	normalizedInput := NormalizeInput(input)

	// Check if the normalized input is empty
	if normalizedInput == "" {
		fmt.Println("Error: Invalid or empty input provided. Please try again.")
		return
	}

	// Load the banner file dynamically
	banner := LoadBanner(font)

	// Generate ASCII art
	asciiArt := GenerateASCIIArt(normalizedInput, banner)

	// If ASCII art is empty, input contained no valid characters
	if len(asciiArt) == 0 {
		fmt.Println("Error: Input contains no valid characters to generate ASCII art.")
		return
	}

	// Display the result
	DisplayASCIIArt(asciiArt)

	fmt.Println("ASCII art generated successfully.")
}
