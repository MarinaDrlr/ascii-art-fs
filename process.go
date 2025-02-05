package main

func GenerateASCIIArt(input string, banner map[rune][]string) []string {
	lines := make([]string, 8) // Preallocate 8 lines for the ASCII output

	for _, char := range input {
		asciiArt, exists := banner[char]
		if !exists {
			asciiArt, exists = banner['?'] // Replace unknown characters with '?'
			if !exists {
				continue // Skip unsupported characters
			}
		}

		for i := range asciiArt {
			lines[i] += asciiArt[i] + " " // Append each line of the ASCII art
		}
	}

	return lines
}
