// process.go
package main

// Generate ASCII art for the input string
func generateASCIIArt(input string) []string {
	lines := make([]string, 8) // Each character's ASCII art is 8 lines tall

	for _, char := range input {
		asciiArt, exists := MapToASCII(char) // Updated to use the exported function
		if !exists {
			lines[0] += "?" // Placeholder for unsupported characters
			continue
		}
		for i, line := range asciiArt {
			lines[i] += line + " " // Add spacing between characters
		}
	}

	return lines
}
