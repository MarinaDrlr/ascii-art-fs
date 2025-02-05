package main

import "fmt"

func GenerateASCIIArt(input string, banner map[rune][]string) []string {
	// Debugging to verify inputs
	fmt.Printf("Debug: Input to GenerateASCIIArt: %s\n", input)
	fmt.Printf("Debug: Banner Map Size: %d\n", len(banner))
	if len(banner) == 0 {
		fmt.Println("Debug: Banner map is empty!")
	}
	lines := make([]string, 8) // Preallocate 8 lines for the ASCII output

	for _, char := range input {
		asciiArt, exists := banner[char]
		if !exists {
			asciiArt, exists = banner['?'] // Replace unknown characters with '?'
			if !exists {
				continue // Skip unsupported characters
			}
		}

		for i := range asciiArt {
			lines[i] += asciiArt[i] + " " // Append each line of the ASCII art
		}
	}

	return lines
}
