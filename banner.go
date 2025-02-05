package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	// Create a map to store ASCII art for each character
	bannerMap := make(map[rune][]string)

	// Open the banner file
	fileName := font + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	// Buffered reading for efficient file processing
	scanner := bufio.NewScanner(file)

	// Variables for storing each character
	var currentChar rune
	var charLines []string

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Check for a new character marker (empty line signals the end of a character)
		if line == "" {
			if len(charLines) > 0 {
				// Add the completed character to the map
				bannerMap[currentChar] = charLines
				charLines = []string{} // Reset for the next character
			}
			continue
		}

		// Check if the line marks a new character
		if len(line) == 1 {
			currentChar = rune(line[0])
			continue
		}

		// Add the line to the current character's ASCII art
		charLines = append(charLines, line)
	}

	// Add the last character (if any)
	if len(charLines) > 0 {
		bannerMap[currentChar] = charLines
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Failed to read banner file: %v\n", err)
		os.Exit(1)
	}

	return bannerMap
}
