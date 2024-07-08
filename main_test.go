package main

import (
	"os"
	"strings"
	"testing"

	"colors/Lib"
)

func TestAsciiArt(t *testing.T) {
	// Stop function from printing to os.Stdout
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = nil

	// Store output string from Lib.AsciiArt
	output := Lib.AsciiArt("", "", "", "Hello", "Hello", "standard.txt")
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

	// Store output string from Lib.AsciiArt
	output = Lib.AsciiArt("", "", "", "hello", "hello", "shadow.txt")
	// Split output into 8 lines
	lines = strings.Split(output, "\n")
	// Declare expected output
	expectedline = []string{
		"                                 ",
		"_|                _| _|          ",
		"_|_|_|     _|_|   _| _|   _|_|   ",
		"_|    _| _|_|_|_| _| _| _|    _| ",
		"_|    _| _|       _| _| _|    _| ",
		"_|    _|   _|_|_| _| _|   _|_|   ",
		"                                 ",
		"                                 ",
	}
	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedline[i] != lines[i] {
			t.Errorf("result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}

	// Store output string from Lib.AsciiArt
	output = Lib.AsciiArt("", "", "", "hello1", "hello1", "thinkertoy.txt")
	// Split output into 8 lines
	lines = strings.Split(output, "\n")
	// Declare expected output
	expectedline = []string{
		"                       ",
		"o        o o       0   ",
		"|        | |      /|   ",
		"O--o o-o | | o-o o |   ",
		"|  | |-' | | | |   |   ",
		"o  o o-o o o o-o o-o-o ",
		"                       ",
		"                       ",
	}
	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedline[i] != lines[i] {
			t.Errorf("result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}

	// Standard output string from Lib.AsciiArt
	output = Lib.AsciiArt("\033[33m", "", "\033[0m", "Hello", "Hello", "standard.txt")

	// Split output into 8 lines
	lines = strings.Split(output, "\n")

	// Declare expected output
	expectedline = []string{
		"\x1b[33m _    _  \x1b[0m\x1b[33m       \x1b[0m\x1b[33m _  \x1b[0m\x1b[33m _  \x1b[0m\x1b[33m        \x1b[0m",
		"\x1b[33m| |  | | \x1b[0m\x1b[33m       \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m        \x1b[0m",
		"\x1b[33m| |__| | \x1b[0m\x1b[33m  ___  \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m  ___   \x1b[0m",
		"\x1b[33m|  __  | \x1b[0m\x1b[33m / _ \\ \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m / _ \\  \x1b[0m",
		"\x1b[33m| |  | | \x1b[0m\x1b[33m|  __/ \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m| | \x1b[0m\x1b[33m| (_) | \x1b[0m",
		"\x1b[33m|_|  |_| \x1b[0m\x1b[33m \\___| \x1b[0m\x1b[33m|_| \x1b[0m\x1b[33m|_| \x1b[0m\x1b[33m \\___/  \x1b[0m",
		"\x1b[33m         \x1b[0m\x1b[33m       \x1b[0m\x1b[33m    \x1b[0m\x1b[33m    \x1b[0m\x1b[33m        \x1b[0m",
		"\x1b[33m         \x1b[0m\x1b[33m       \x1b[0m\x1b[33m    \x1b[0m\x1b[33m    \x1b[0m\x1b[33m        \x1b[0m",
	}

	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedline[i] != lines[i] {
			t.Errorf("Standard error. Result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("Standard error. Expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}

	// Shadow output test
	output = Lib.AsciiArt("\033[33m", "", "\033[0m", "Hello", "Hello", "shadow.txt")

	// Split output into 8 lines
	lines = strings.Split(output, "\n")

	// Declare expected output
	expectedShadow := []string{
		"\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m   \x1b[0m\x1b[33m   \x1b[0m\x1b[33m         \x1b[0m",
		"\x1b[33m_|    _| \x1b[0m\x1b[33m         \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m         \x1b[0m",
		"\x1b[33m_|    _| \x1b[0m\x1b[33m  _|_|   \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m  _|_|   \x1b[0m",
		"\x1b[33m_|_|_|_| \x1b[0m\x1b[33m_|_|_|_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_|    _| \x1b[0m",
		"\x1b[33m_|    _| \x1b[0m\x1b[33m_|       \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_|    _| \x1b[0m",
		"\x1b[33m_|    _| \x1b[0m\x1b[33m  _|_|_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m_| \x1b[0m\x1b[33m  _|_|   \x1b[0m",
		"\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m   \x1b[0m\x1b[33m   \x1b[0m\x1b[33m         \x1b[0m",
		"\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m   \x1b[0m\x1b[33m   \x1b[0m\x1b[33m         \x1b[0m",
	}

	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedShadow[i] != lines[i] {
			t.Errorf("Shadow error. Result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("Shadow error. Expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}

	// THinkertoy output test
	output = Lib.AsciiArt("\033[33m", "", "\033[0m", "Hello", "Hello", "thinkertoy.txt")

	// Split output into 8 lines
	lines = strings.Split(output, "\n")

	// Declare expected output
	expectedThinkertoy := []string{
		"\x1b[33m     \x1b[0m\x1b[33m    \x1b[0m\x1b[33m  \x1b[0m\x1b[33m  \x1b[0m\x1b[33m    \x1b[0m",
		"\x1b[33mo  o \x1b[0m\x1b[33m    \x1b[0m\x1b[33mo \x1b[0m\x1b[33mo \x1b[0m\x1b[33m    \x1b[0m",
		"\x1b[33m|  | \x1b[0m\x1b[33m    \x1b[0m\x1b[33m| \x1b[0m\x1b[33m| \x1b[0m\x1b[33m    \x1b[0m",
		"\x1b[33mO--O \x1b[0m\x1b[33mo-o \x1b[0m\x1b[33m| \x1b[0m\x1b[33m| \x1b[0m\x1b[33mo-o \x1b[0m",
		"\x1b[33m|  | \x1b[0m\x1b[33m|-' \x1b[0m\x1b[33m| \x1b[0m\x1b[33m| \x1b[0m\x1b[33m| | \x1b[0m",
		"\x1b[33mo  o \x1b[0m\x1b[33mo-o \x1b[0m\x1b[33mo \x1b[0m\x1b[33mo \x1b[0m\x1b[33mo-o \x1b[0m",
		"\x1b[33m     \x1b[0m\x1b[33m    \x1b[0m\x1b[33m  \x1b[0m\x1b[33m  \x1b[0m\x1b[33m    \x1b[0m",
		"\x1b[33m     \x1b[0m\x1b[33m    \x1b[0m\x1b[33m  \x1b[0m\x1b[33m  \x1b[0m\x1b[33m    \x1b[0m",
	}

	// Compare expected and output
	for i := 0; i < 8; i++ {
		if expectedThinkertoy[i] != lines[i] {
			t.Errorf("Thinkertoy error. Result: %q, Len: %d\n", lines[i], len(lines[i]))
			t.Errorf("Thinkertoy error. Expect: %q, Len: %d\n", expectedline[i], len(expectedline[i]))
			t.Errorf("Test Failed!")
		}
	}
}

func TestEscapeSequence(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"first test", "\\t", true},
		{"second test", "\\n", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Lib.EscapeSequence(test.input); got != test.want {
				t.Errorf("EscapeSequence(%s) = %v, want %v", test.name, got, test.want)
			}
		})
	}
}
