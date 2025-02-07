package main

import "strings"

// NormalizeInput replaces the literal `\n` (backslash + n) with actual newlines
func NormalizeInput(input string) string {
	return strings.ReplaceAll(input, `\n`, "\n")
}
