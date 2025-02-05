package main

import "fmt"

func GenerateASCIIArt(input string, banner map[rune][]string) []string {
	lines := make([]string, 8) // Each character's ASCII art is 8 lines tall

	fmt.Printf("Debug: Input to GenerateASCIIArt: %s\n", input) // Log the input

	for _, char := range input {
		if char == '\n' {
			// Handle newline in the input
			fmt.Println("Debug: Encountered newline character. Skipping to next line.")
			for i := range lines {
				lines[i] += "\n"
			}
			continue
		}

		asciiArt, exists := banner[char]
		if !exists {
			// If the character is not found in the banner map
			fmt.Printf("Warning: Character '%c' not found in banner map.\n", char)
			continue
		}

		fmt.Printf("Debug: Adding ASCII art for character '%c' to lines.\n", char)
		for i, artLine := range asciiArt {
			lines[i] += artLine + " " // Add spacing between characters
		}
	}

	fmt.Printf("Debug: Final ASCII art lines: %v\n", lines) // Log the result
	return lines
}
