package main

import (
	"bufio"
	"fmt"
	"os"
)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadBanner loads the banner file and constructs a map of characters to their ASCII art representation
func LoadBanner(font string) map[rune][]string {
	bannerMap := make(map[rune][]string) // Map to store the banner characters
	var currentChar rune                 // Variable to hold the currently processed character
	charLines := []string{}              // Slice to store lines of ASCII art for a character

	// Open the banner file
	file, err := os.Open(font + ".txt")
	if err != nil {
		fmt.Printf("Error: Could not open banner file: %s\n", font+".txt")
		os.Exit(1)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line from the file

		// Debugging: Show the current line being processed
		fmt.Printf("Debug: Current line read: %q\n", line)

		// Check for an empty line, signaling the end of a character block
		if line == "" {
			if len(charLines) == 8 {
				// Add the character and its ASCII art to the map
				bannerMap[currentChar] = charLines
				fmt.Printf("Debug: Completed character '%c' with lines: %v\n", currentChar, charLines)
			} else if len(charLines) > 0 {
				// Warn if a block is incomplete
				fmt.Printf("Warning: Incomplete block for character '%c', lines: %v\n", currentChar, charLines)
			}
			// Reset for the next character
			charLines = []string{}
			continue
		}

		// Check for a new character declaration (single character line)
		if len(charLines) == 0 && len(strings.TrimSpace(line)) == 1 {
			// Detect character declaration
			currentChar = rune(line[0])
			fmt.Printf("Debug: Detected new character: '%c'\n", currentChar)
		} else if len(charLines) < 8 {
			// Collect ASCII art lines (up to 8 per character)
			charLines = append(charLines, line)
			fmt.Printf("Debug: Adding line to character '%c': %q\n", currentChar, line)
		} else {
			// Unexpected extra lines (beyond 8)
			fmt.Printf("Warning: Unexpected extra line for character '%c': %q\n", currentChar, line)
		}
	}

	// Add the last character block if not followed by an empty line
	if len(charLines) == 8 {
		bannerMap[currentChar] = charLines
		fmt.Printf("Debug: Added last character '%c' to banner map with lines: %v\n", currentChar, charLines)
	} else if len(charLines) > 0 {
		// Warn about incomplete block at the end of the file
		fmt.Printf("Warning: Last character '%c' block incomplete with %d lines: %v\n", currentChar, len(charLines), charLines)
	}

	// Check for errors in scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Failed to read the banner file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Debug: Banner map size: %d, keys: %v\n", len(bannerMap), getKeys(bannerMap))
	return bannerMap
}

// Helper function to get the keys of the banner map for debugging
func getKeys(m map[rune][]string) []rune {
	keys := []rune{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

