package Lib

import (
	"fmt"
	"strings"
)

// Prints new lines and calls HandleCharacters() on each word
func HandleWords(color string, reset string, subString string, slices []string, words []string) string {
	var output string
	countSpaces := 0

	for _, word := range words {
		if word == "" {
			countSpaces++
			if countSpaces < len(words) {
				fmt.Println()
			}
		} else {
			output = HandleCharacters(color, reset, subString, word, slices)
		}
	}
	return output
}

// Prints the ASCII ART for each character in each word and updates the output string
func HandleCharacters(color string, reset string, subString string, word string, slices []string) string {
	// Find the matching characters
	match, _ := matchingStr(word, subString)

	// Initialize utility variables
	var startIndex int
	output := ""

	// Loop through each line of the ASCII art to be printed (8 lines)
	for i := 0; i < 8; i++ {

		// Loop through each character in the word
		for _, ch := range word {

			// Calculate the index of the first line of art for each character
			startIndex = int(ch-32)*9 + 1

			// If the character matches the substring, add colored ASCII art to the output
			if match[string(ch)] {
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
