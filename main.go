package main

import "fmt"

func main() {
	input := getInput()
	normalizedInput := normalizeInput(input)
	asciiArt := generateASCIIArt(normalizedInput)
	displayASCIIArt(asciiArt)
	fmt.Println("ASCII art generated successfully.")
}
