package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(font string) map[rune][]string {
	filename := font + ".txt" // Choose the correct file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: Could not open banner file:", filename)
		os.Exit(1)
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var currentChar rune = ' ' // Start with space character
	charLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line signals end of a character block
		if line == "" {
			if len(charLines) > 0 {
				bannerMap[currentChar] = charLines
				currentChar++
				charLines = []string{} // Reset for the next character
			}
			continue
		}

		charLines = append(charLines, line)
	}

	// Add the last character block (if not followed by an empty line)
	if len(charLines) > 0 {
		bannerMap[currentChar] = charLines
	}

	return bannerMap
}
