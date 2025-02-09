package funcs

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner reads the font file and maps characters to ASCII art
func LoadBanner(font string) (map[rune][]string, error) {
	filename := "../fonts/" + font + ".txt"

	// fmt.Println("Looking for font file at:", filename)

	// Check if the file exists before trying to open it
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: Banner font \"%s\" does not exist.", font)
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error: Could not open banner file \"%s\".", filename)
	}
	defer file.Close()

	bannerMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var currentChar rune = ' ' // Start with space character
	var charLines []string
	linesRead := 0 // Counter to check if file is empty

	for scanner.Scan() {
		linesRead++ // Increment counter for each line read
		line := scanner.Text()

		if line == "" {
			if len(charLines) > 0 {
				bannerMap[currentChar] = append([]string{}, charLines...) // Ensure a copy is stored
				currentChar++
				charLines = nil
			}
			continue
		}

		charLines = append(charLines, line)
	}

	// If no lines were read, return an error
	if linesRead == 0 {
		return nil, fmt.Errorf("Error: Banner file \"%s\" is empty.", filename)
	}

	if len(charLines) > 0 {
		bannerMap[currentChar] = append([]string{}, charLines...)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error: Failed to read banner file \"%s\": %s", filename, err)
	}

	return bannerMap, nil
}
