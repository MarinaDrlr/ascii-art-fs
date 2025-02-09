package main

import (
	"ascii-art/funcs"
	"testing"
)

func TestNormalizeInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
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

func TestLoadBanner(t *testing.T) {
	banner, err := funcs.LoadBanner("standard")
	if err != nil {
		t.Errorf("LoadBanner failed: %v", err)
	}

	if len(banner) == 0 {
		t.Errorf("LoadBanner failed: expected non-empty banner map")
	}
}

func TestGenerateASCIIArt(t *testing.T) {
	banner, err := funcs.LoadBanner("standard")
	if err != nil {
		t.Errorf("LoadBanner failed: %v", err)
	}

	tests := []struct {
		input    string
		expected int
	}{
		{"A", 8},
		{"Hello", 8},
		{"Hello\nWorld", 16},
	}

	for _, test := range tests {
		result, err := funcs.GenerateASCIIArt(test.input, banner)
		if err != nil {
			t.Errorf("GenerateASCIIArt failed for input %q: %v", test.input, err)
		}

		if len(result) != test.expected {
			t.Errorf("GenerateASCIIArt(%q) returned %d lines; want %d", test.input, len(result), test.expected)
		}
	}
}
