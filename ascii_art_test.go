package main

import (
	"ascii-art_1/funcs"
	"testing"
)

// TestNormalizeInput checks if NormalizeInput correctly replaces `\n` (literal backslash + n) with actual newlines
func TestNormalizeInput(t *testing.T) {
	tests := []struct {
		input    string // Input string to be tested
		expected string // Expected output after normalization
	}{
		{`Hello\nWorld`, "Hello\nWorld"},
		{`Line1\nLine2\nLine3`, "Line1\nLine2\nLine3"},
	}

	for _, test := range tests {
		result := funcs.NormalizeInput(test.input)
		if result != test.expected {
			t.Errorf("NormalizeInput(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// TestLoadBanner ensures that the banner font is loaded correctly and is not empty
func TestLoadBanner(t *testing.T) {
	banner, err := funcs.LoadBanner("standard")
	if err != nil {
		t.Errorf("LoadBanner failed: %v", err)
	}

	// The banner map should not be empty if the font loaded successfully
	if len(banner) == 0 {
		t.Errorf("LoadBanner failed: expected non-empty banner map")
	}
}

// TestGenerateASCIIArt verifies that ASCII art is generated correctly for different inputs
func TestGenerateASCIIArt(t *testing.T) {
	// Load the standard font banner for testing
	banner, err := funcs.LoadBanner("standard")
	if err != nil {
		t.Errorf("LoadBanner failed: %v", err)
	}

	tests := []struct {
		input    string // Input text to convert to ASCII art
		expected int    // Expected number of lines in the ASCII output
	}{
		{"A", 8},             // A single character should have 8 lines
		{"Hello", 8},         // A word should also be 8 lines tall
		{"Hello\nWorld", 16}, // Two lines should result in 16 total lines
	}

	for _, test := range tests {
		result, err := funcs.GenerateASCIIArt(test.input, banner)
		if err != nil {
			t.Errorf("GenerateASCIIArt failed for input %q: %v", test.input, err)
		}

		// Ensure the number of lines in the result matches the expected value
		if len(result) != test.expected {
			t.Errorf("GenerateASCIIArt(%q) returned %d lines; want %d", test.input, len(result), test.expected)
		}
	}
}
