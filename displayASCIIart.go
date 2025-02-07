package main

import "fmt"

// DisplayASCIIArt prints each line of the ASCII art
func DisplayASCIIArt(asciiArt []string) {
	for _, line := range asciiArt {
		fmt.Println(line)
	}
}
