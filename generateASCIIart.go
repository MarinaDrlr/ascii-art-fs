package main

import (
	"fmt"
)

// GenerateASCIIArt converts input into ASCII art based on the loaded banner
func GenerateASCIIArt(input string, banner map[rune][]string) ([]string, error) {
	var result []string
	lines := make([]string, 8)  // Each character's ASCII art is 8 lines tall
	newLineEncountered := false // Track if we processed a newline

	for _, char := range input {
		if char == '\n' {
			// If there's content in lines, flush it to the result
			if len(lines[0]) > 0 {
				result = append(result, lines...)
				lines = make([]string, 8) // Reset lines
			}

			// Only add a blank line if a previous newline was already encountered
			if newLineEncountered {
				result = append(result, "")
			}

			newLineEncountered = true
			continue
		}

		// Reset the newline tracker because we found a character
		newLineEncountered = false

		asciiArt, exists := banner[char]
		if !exists {
			return nil, fmt.Errorf("Error: Character '%c' not found in banner file", char)
		}

		// Ensure the ASCII art has exactly 8 lines to prevent index out of range errors
		if len(asciiArt) != 8 {
			return nil, fmt.Errorf("Error: Character '%c' has an invalid ASCII art format", char)
		}

		for i, artLine := range asciiArt {
			lines[i] += artLine + " " // Add spacing between characters
		}
	}

	// Append the last processed text if it's not empty
	if len(lines[0]) > 0 {
		result = append(result, lines...)
	}

	// Ensure the output ends with a blank line if the input ended with '\n'
	if len(input) > 0 && input[len(input)-1] == '\n' {
		result = append(result, "")
	}

	return result, nil
}
