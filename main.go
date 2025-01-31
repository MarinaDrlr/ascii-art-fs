package main

import "fmt"

func main() {
	// Get user input and font choice
	input, font := GetInput()

	// Load the banner file dynamically
	banner := LoadBanner(font)

	// Generate ASCII art
	asciiArt := GenerateASCIIArt(input, banner)

	// Display the result
	DisplayASCIIArt(asciiArt)

	fmt.Println("ASCII art generated successfully.")
}
