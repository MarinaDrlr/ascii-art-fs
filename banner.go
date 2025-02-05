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

	// Buffered reading
	scanner := bufio.NewScanner(file)

	// Variables for processing
	var currentChar rune
	var charLines []string
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Debug: Log each line as it is read
		fmt.Printf("Debug: Read line: %s\n", line)

		// Handle empty line (end of character block)
		if line == "" {
			if len(charLines) > 0 {
				fmt.Printf("Debug: Adding character '%c' to banner map with %d lines\n", currentChar, len(charLines))
				bannerMap[currentChar] = charLines
				charLines = []string{}
				lineCount = 0
			}
			continue
		}

		// Detect new character
		if lineCount == 0 {
			currentChar = rune(line[0]) // First character in the line
			fmt.Printf("Debug: Detected new character: %c\n", currentChar)
			lineCount++
			continue
		}

		// Collect ASCII art for the current character
		charLines = append(charLines, line)
		lineCount++
	}

	// Add the last character block if it exists
	if len(charLines) > 0 {
		fmt.Printf("Debug: Adding last character '%c' to banner map with %d lines\n", currentChar, len(charLines))
		bannerMap[currentChar] = charLines
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Failed to read banner file: %v\n", err)
		os.Exit(1)
	}

	return bannerMap
}
