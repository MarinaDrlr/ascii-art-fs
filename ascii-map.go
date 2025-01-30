// ascii_map.go
package main

// Mock ASCII Map for demonstration purposes
var asciiMap = map[rune][]string{
	'A': {"  A  ", " A A ", "AAAAA", "A   A", "A   A", "     ", "     ", "     "},
	'B': {"BBBB ", "B   B", "BBBB ", "B   B", "BBBB ", "     ", "     ", "     "},
	' ': {"     ", "     ", "     ", "     ", "     ", "     ", "     ", "     "},
}

// MapToASCII maps a character to its ASCII art representation
func MapToASCII(char rune) ([]string, bool) {
	art, exists := asciiMap[char]
	return art, exists
}
