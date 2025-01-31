package main

import "fmt"

func main() {
	// Get user input and font choice
	input, font := GetInput()
	// Normalize the input to match the banner map keys
	normalizedInput := NormalizeInput(input)

	// Load the banner file dynamically
	banner := LoadBanner(font)

	// Generate ASCII art
	asciiArt := GenerateASCIIArt(normalizedInput, banner)

	// Display the result
	DisplayASCIIArt(asciiArt)

	fmt.Println("ASCII art generated successfully.")
}
