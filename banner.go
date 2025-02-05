package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	// Create a map to store the ASCII art for each character
	bannerMap := make(map[rune][]string)

	// Open the banner file
	fileName := font + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	// Debug: Notify the start of file reading
	fmt.Printf("Debug: Starting to load banner from %s\n", fileName)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var charLines []string
	var currentChar rune = ' ' // Placeholder for the current character

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Debug: Current line read: %q\n", line) // Debug: Print each line

		// Check for an empty line, which indicates the end of a character block
		if line == "" {
			if len(charLines) == 8 {
				// Add the character and its ASCII art to the map
				bannerMap[currentChar] = charLines
				fmt.Printf("Debug: Completed character '%c' with lines: %v\n", currentChar, charLines)
			} else if len(charLines) > 0 {
				// Debug: Warn if a block is incomplete
				fmt.Printf("Warning: Incomplete block for character '%c', lines: %v\n", currentChar, charLines)
			}
			// Reset for the next character
			charLines = []string{}
			continue
		}

		// Check for a new character declaration
		if len(charLines) == 0 {
			// Assume the line is the ASCII character declaration
			if len(line) == 1 {
				currentChar = rune(line[0])
				fmt.Printf("Debug: Detected new character: '%c'\n", currentChar)
			} else {
				fmt.Printf("Warning: Unexpected character line: %q\n", line)
			}
		} else {
			// Add the line to the current character block
			charLines = append(charLines, line)
		}
	}

	// Add the last character block (if not followed by an empty line)
	if len(charLines) == 8 {
		bannerMap[currentChar] = charLines
		fmt.Printf("Debug: Adding last character '%c' to banner map with lines: %v\n", currentChar, charLines)
	} else if len(charLines) > 0 {
		fmt.Printf("Warning: Last character '%c' block incomplete with %d lines\n", currentChar, len(charLines))
	}

	// Debug: Print banner map size and keys
	fmt.Printf("Debug: Banner map size: %d, keys: %v\n", len(bannerMap), keys(bannerMap))

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Failed to read banner file: %s\n", fileName)
		os.Exit(1)
	}

	return bannerMap
}

// Helper function to get keys from a map
func keys(m map[rune][]string) []rune {
	k := make([]rune, 0, len(m))
	for key := range m {
		k = append(k, key)
	}
	return k
}
