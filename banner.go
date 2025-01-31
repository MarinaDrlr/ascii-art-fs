package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner reads the banner file and stores characters in a map
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
			bannerMap[currentChar] = charLines
			currentChar++
			charLines = []string{} // Reset for the next character
			continue
		}

		charLines = append(charLines, line)
	}

	return bannerMap
}
