package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	bannerMap := make(map[rune][]string)
	fileName := font + ".txt"

	// Open the banner file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentChar rune
	var charLines []string
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Debugging: Print each line read
		fmt.Printf("Debug: Read line: %q\n", line)

		// Handle empty lines as separator
		if line == "" {
			if len(charLines) == 8 {
				bannerMap[currentChar] = charLines
				fmt.Printf("Debug: Added character '%c' to banner map with %d lines\n", currentChar, len(charLines))
			} else if len(charLines) > 0 {
				fmt.Printf("Warning: Character '%c' block incomplete with %d lines\n", currentChar, len(charLines))
			}
			// Reset for the next character block
			charLines = []string{}
			lineCount = 0
			continue
		}

		// First non-empty line marks the character
		if lineCount == 0 {
			currentChar = rune(line[0]) // Extract character
			fmt.Printf("Debug: Detected new character: '%c'\n", currentChar)
			lineCount++
			continue
		}

		// Append line to character block
		charLines = append(charLines, line)
		lineCount++

		// If block completes 8 lines, store it
		if len(charLines) == 8 {
			bannerMap[currentChar] = charLines
			fmt.Printf("Debug: Added character '%c' to banner map with %d lines (complete block)\n", currentChar, len(charLines))
			// Reset for the next character block
			charLines = []string{}
			lineCount = 0
		}
	}

	// Handle the last character block if file doesn't end with an empty line
	if len(charLines) == 8 {
		bannerMap[currentChar] = charLines
		fmt.Printf("Debug: Added last character '%c' to banner map with %d lines\n", currentChar, len(charLines))
	} else if len(charLines) > 0 {
		fmt.Printf("Warning: Last character '%c' block incomplete with %d lines\n", currentChar, len(charLines))
	}

	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Failed to read banner file: %v\n", err)
		os.Exit(1)
	}

	return bannerMap
}
