package Lib

import (
	"fmt"
	"os"
	"strings"
)

// Prints new lines and calls HandleCharacters() on each word
func HandleWords(color, reset string, subStrings, slices, words []string) string {
	var start, end int = -1, -1
	var output string
	countSpaces := 0

	// Print one fewer newline for every empty string
	for _, word := range words {
		if word == "" {
			countSpaces++
			if countSpaces < len(words) {
				fmt.Println()
			}
		} else {
			// Find range of matching substring
			var x, y int
			found, x, y := subToColor(word, subStrings)

			// Store the values of start and end when found is true
			if found {
				start = x
				end = y
			}
			output = HandleCharacters(start, end, color, reset, word, subStrings, slices)
			start, end = -1, -1
		}
	}
	return output
}

// Prints the ASCII ART for each character in each word and updates the output string
func HandleCharacters(start, end int, color, reset, word string, subStrings, slices []string) string {
	// Initialize utility variables
	var startIndex int
	output := ""

	// Loop through each line of the ASCII art to be printed (8 lines)
	for i := 0; i < 8; i++ {

		// Loop through each character in the word
		for j, ch := range word {
			// Calculate the index of the first line of art for each character
			startIndex = int(ch-32)*9 + 1

			// If j is within start and stop color the output and reset
			if j >= start && j <= end {
				output += color + slices[startIndex+i] + reset
				fmt.Printf("%s%s%s", color, slices[startIndex+i], reset)
			} else {
				// If the character doesn't match, add regular ASCII art to the output
				output += slices[startIndex+i]
				fmt.Printf("%s", slices[startIndex+i])
			}
		}

		// Add a new line after each line for testing
		output += "\n"
		fmt.Println()
	}

	// Return the updated output string
	return output
}

// Chaeck if input contains non-printable ascii characters with escape sequences (except '\n')
func EscapeSequence(input string) bool {
	allowedEscapes := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r"}
	for _, unprint := range allowedEscapes {
		// Detect escape sequence characters in input string
		if strings.Contains(input, unprint) {
			return true
		}
	}
	return false
}

// Check if input contains characters outside of the banner file (except '\n')
func IsPrintable(input string) bool {
	var status bool

	// Check if character is represented in banner file
	for _, v := range input {
		if (v >= ' ' && v <= '~') || v == '\n' {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
}

// Prints the standard error message
func PrintError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	os.Exit(0)
}

// Function to trim whitespaces at extreme ends
func trimSpace(s string) string {
	if s == "" {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == ' ' {
			s = s[1:]
			s = trimSpace(s)
		} else if i == len(s)-1 && s[i] == ' ' {
			s = s[:len(s)-1]
			s = trimSpace(s)
		}
	}
	return s
}

// Returns range of substring to be colored and true if a match is found
func subToColor(s string, subStrings []string) (bool, int, int) {
	for _, sub := range subStrings {
		for i := 0; i < len(s); i++ {
			if i <= len(s)-len(sub) && s[i:i+len(sub)] == sub {
				return true, i, i + len(sub) - 1
			}
		}
	}

	return false, 0, 0
}
