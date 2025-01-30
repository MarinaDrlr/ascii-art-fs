package main

import "fmt"

func main() {
	input := GetInput()
	normalizedInput := NormalizeInput(input)
	asciiArt := generateASCIIArt(normalizedInput)
	displayASCIIArt(asciiArt)
	fmt.Println("ASCII art generated successfully.")
}
