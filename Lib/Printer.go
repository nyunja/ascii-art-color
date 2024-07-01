package Lib

import (
	"fmt"
	"os"
)

// Prints new lines and calls HandleCharacters() on each word
func HandleWords(color1, color2, reset string, subStrings, slices, words []string) string {
	var start, end int = -1, -1
	var output string
	countSpaces := 0

	// Print one fewer newline for every empty string
	for _, word := range words {
		if word == "" {
			countSpaces++
			if countSpaces < len(words) {
				output += "\n"
			}

			// Print Ascii Art
		} else {
			output += HandleCharacters(start, end, color1, color2, reset, word, subStrings, slices)
			start, end = -1, -1
		}
	}
	return output
}

// Prints the ASCII ART for each character in each word and updates the output string
func HandleCharacters(start, end int, color1, color2, reset, word string, subStrings, slices []string) string {
	// Initialize utility variables

	var startIndex int
	output := ""

	// Loop through each line of the ASCII art to be printed (8 lines)
	for i := 0; i < 8; i++ {

		// Loop through each character in the word
		for j, ch := range word {
			// Find range of matching substring
			found, x, y := subToColor(j, word, subStrings)

			// Store the values of start and end when found is true
			if found {
				start = x
				end = y
			}
			// Calculate the index of the first line of art for each character
			startIndex = int(ch-32)*9 + 1

			// If j is within start and stop color the output and reset
			if len(color1) > 0 && j >= start && j <= end {
				output += color1 + slices[startIndex+i] + reset

				// Print Ascii Art with second color
			} else if len(color2) > 0 {
				output += color2 + slices[startIndex+i] + reset

				// If the character doesn't match, add regular ASCII art to the output
			} else {
				output += slices[startIndex+i]
			}
		}

		// Add a new line after each line for testing
		output += "\n"
	}

	// Return the updated output string
	return output
}

// Prints the standard error message
func PrintError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	os.Exit(0)
}
