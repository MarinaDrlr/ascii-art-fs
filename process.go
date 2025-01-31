package main

// GenerateASCIIArt constructs ASCII representation using the loaded banner
func GenerateASCIIArt(input string, banner map[rune][]string) []string {
	lines := make([]string, 8) // ASCII characters are 8 lines tall

	for _, char := range input {
		asciiArt, exists := banner[char]
		if !exists {
			continue // Skip unknown characters
		}

		for i, line := range asciiArt {
			lines[i] += line + " "
		}
	}

	return lines
}
