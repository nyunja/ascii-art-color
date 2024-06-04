package main

import (
	"colors/Lib"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	// Stop function from printing to os.Stdout
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = nil

	// Store output string from Lib.AsciiArt
	output := Lib.AsciiArt("\033[37m", "\033[0m", "Hello", "Hello", "standard.txt")

	// Split output into 8 lines
	lines := strings.Split(output, "\n")

	// Declare expected output
	expectedline := []string{
		"\x1b[37m _    _  \x1b[0m\x1b[37m       \x1b[0m\x1b[37m _  \x1b[0m\x1b[37m _  \x1b[0m\x1b[37m        \x1b[0m",
		"\x1b[37m| |  | | \x1b[0m\x1b[37m       \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m        \x1b[0m",
		"\x1b[37m| |__| | \x1b[0m\x1b[37m  ___  \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m  ___   \x1b[0m",
		"\x1b[37m|  __  | \x1b[0m\x1b[37m / _ \\ \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m / _ \\  \x1b[0m",
		"\x1b[37m| |  | | \x1b[0m\x1b[37m|  __/ \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m| | \x1b[0m\x1b[37m| (_) | \x1b[0m",
		"\x1b[37m|_|  |_| \x1b[0m\x1b[37m \\___| \x1b[0m\x1b[37m|_| \x1b[0m\x1b[37m|_| \x1b[0m\x1b[37m \\___/  \x1b[0m",
		"\x1b[37m         \x1b[0m\x1b[37m       \x1b[0m\x1b[37m    \x1b[0m\x1b[37m    \x1b[0m\x1b[37m        \x1b[0m",
		"\x1b[37m         \x1b[0m\x1b[37m       \x1b[0m\x1b[37m    \x1b[0m\x1b[37m    \x1b[0m\x1b[37m        \x1b[0m",
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
