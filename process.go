package main

import "strings"

func GenerateASCIIArt(input string, banner map[rune][]string) []string {
	lines := []string{}                      // To store the resulting ASCII art
	inputLines := strings.Split(input, "\n") // Split input into separate lines

	for _, line := range inputLines {
		asciiLines := make([]string, 8) // ASCII characters are 8 lines tall

		for _, char := range line {
			asciiArt, exists := banner[char]
			if !exists {
				asciiArt, exists = banner['?'] // Replace unknown characters with '?'
				if !exists {
					continue // Skip if '?' isn't in the banner file either
				}
			}

			for i, artLine := range asciiArt {
				asciiLines[i] += artLine + " "
			}
		}

		// Add the generated ASCII art for the line to the result
		lines = append(lines, asciiLines...)
		lines = append(lines, "") // Add a blank line between each text line
	}

	return lines
}
