package main

import "fmt"

// Display the ASCII art by printing each line
func displayASCIIArt(asciiArt []string) {
	for _, line := range asciiArt {
		fmt.Println(line)
	}
}
