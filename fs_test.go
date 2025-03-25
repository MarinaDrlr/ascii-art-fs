package main

import (
	"fs/funcs"
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

// TestNormalizeInputNewlines checks that actual newlines are preserved
func TestNormalizeInputNewlines(t *testing.T) {
	input := "Line1\nLine2"
	expected := "Line1\nLine2"

	result := funcs.NormalizeInput(input)
	if result != expected {
		t.Errorf("NormalizeInput with real newlines returned %q; want %q", result, expected)
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

// TestLoadCorruptedBanner checks that an error is returned
// when trying to load a corrupted or incomplete banner file
func TestLoadCorruptedBanner(t *testing.T) {
	// "broken.txt" should be a corrupted banner file in the "fonts" directory
	_, err := funcs.LoadBanner("broken")
	if err == nil {
		t.Errorf("Expected error when loading corrupted banner, got nil")
	}
}

// TestLoadMissingFontFile checks that an error is returned
// when trying to load a banner font file that does not exist
func TestLoadMissingFontFile(t *testing.T) {
	_, err := funcs.LoadBanner("doesnotexist")
	if err == nil {
		t.Errorf("Expected error when loading missing font file, got nil")
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

// TestGenerateASCIIArtUnsupportedChar checks that an error is returned
// when the input contains a character not found in the banner map
func TestGenerateASCIIArtUnsupportedChar(t *testing.T) {
	// Load the standard banner
	banner, err := funcs.LoadBanner("standard")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	// Input contains a non-ASCII character (emoji)
	_, err = funcs.GenerateASCIIArt("Hello 🌍", banner)
	if err == nil {
		t.Errorf("Expected error for unsupported character, got nil")
	}
}

// // TestDEBUGCheck forces a failure to ensure tests are running
// func TestDEBUGCheck(t *testing.T) {
// 	t.Fail() // This will make the test fail on purpose
// }
