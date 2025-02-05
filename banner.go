package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	bannerMap := make(map[rune][]string)
	fileName := font + ".txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	// Buffered file reading
	scanner := bufio.NewScanner(file)

	// Variables to process characters
	var currentChar rune
	var charLines []string
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// If an empty line is detected
		if line == "" {
			if len(charLines) > 0 {
				bannerMap[currentChar] = charLines // Add completed character
				charLines = []string{}             // Reset for the next character
				lineCount = 0                      // Reset line count
			}
			continue
		}

		// Check if the line defines a new character (first line after empty line)
		if lineCount == 0 {
			currentChar = rune(line[0]) // First character in the line is the key
			lineCount++                 // Increment line count
			continue
		}

		// Add the ASCII art line for the current character
		charLines = append(charLines, line)
		lineCount++
	}

	// Add the last character if it exists
	if len(charLines) > 0 {
		bannerMap[currentChar] = charLines
	}

	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading banner file: %v\n", err)
		os.Exit(1)
	}

	return bannerMap
}
