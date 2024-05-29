package main

import (
	"colors/Lib"
	"os"
	"strings"
	"testing"
)

// func TestHasSuffix(t *testing.T) {
// 	expected := true
// 	result := hasSuffix("hello", "o")
// 	if result != expected {
// 		t.Errorf("Try again")
// 	}
// }

// func TestHasPrefix(t *testing.T) {
// 	expected := true
// 	result := hasPrefix("hellored", "hello")
// 	if result != expected {
// 		t.Errorf("Try again")
// 	}
// }

func TestAsciiArt(t *testing.T) {
	// Stop function from printing to os.Stdout
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = nil

	// Store output string from Lib.AsciiArt
	output := Lib.AsciiArt("","","","Hello", "")

	// Split output into 8 lines
	lines := strings.Split(output, "\n")

	// Declare expected output
	expectedline := []string{
		" _    _          _   _          ",
		"| |  | |        | | | |         ",
		"| |__| |   ___  | | | |   ___   ",
		"|  __  |  / _ \\ | | | |  / _ \\  ",
		"| |  | | |  __/ | | | | | (_) | ",
		"|_|  |_|  \\___| |_| |_|  \\___/  ",
		"                                ",
		"                                ",
	}

	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedline[i] != lines[i] {
			t.Errorf("result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}
}
