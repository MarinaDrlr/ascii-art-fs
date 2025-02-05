package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	bannerMap := make(map[rune][]string)

	// Open the banner file
	file, err := os.Open(font + ".txt")
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", font+".txt")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentChar rune
	var charLines []string

	for scanner.Scan() {
		line := scanner.Text()

		// Check for an empty line indicating the end of a character block
		if line == "" {
			if len(charLines) == 8 {
				// Store the character and its ASCII art
				bannerMap[currentChar] = charLines
			}
			// Reset for the next character
			charLines = []string{}
			continue
		}

		// Detect new character declaration (single character line)
		if len(charLines) == 0 && len(line) == 1 {
			currentChar = rune(line[0])
		} else if len(charLines) < 8 {
			// Add the line to the current character's lines
			charLines = append(charLines, line)
		}
	}

	// Add the last character block if complete
	if len(charLines) == 8 {
		bannerMap[currentChar] = charLines
	}

	// Check for errors during file scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Could not read banner file: %s\n", err)
		os.Exit(1)
	}

	return bannerMap
}
