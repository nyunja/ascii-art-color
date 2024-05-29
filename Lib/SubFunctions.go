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
	match, _ := matchingStr(word, subString)
	// if !found {
	// 	fmt.Printf("Error: %s has no match\n", subString)
	// 	os.Exit(0)
	// }
	var startIndex int
	output := ""
	for i := 0; i < 8; i++ {
		for _, ch := range word {
			startIndex = int(ch-32)*9 + 1
			if match[string(ch)] {
				output += color + slices[startIndex+i] + reset
				fmt.Printf("%s%s%s", color, slices[startIndex+i], reset)
			} else {
				output += slices[startIndex+i]
				fmt.Printf("%s", slices[startIndex+i])
				// else if matchingStr(subString, string(ch)) && match[string(ch)] {
				// 	fmt.Printf("%s%s%s", color, string(ch), reset)
				// } else if !matchingStr(subString, string(ch)) && !match[string(ch)] {
				// 	fmt.Printf("%s", string(ch))
				// }
			}
			// startIndex = int(ch-32)*9 + 1
			// output += slices[startIndex+i]
			// fmt.Printf("%s", slices[startIndex+i])
		}
		output += "\n"
		fmt.Println()
	}
	return output
}

// Chaeck if input contains non-printable ascii characters with escape sequences (except '\n')
func EscapeSequence(input string) bool {
	allowedEscapes := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r"}
	for _, unprint := range allowedEscapes {
		if strings.Contains(input, unprint) {
			return true
		}
	}
	return false
}

// Check if input contains characters outside of the banner file (except '\n')
func IsPrintable(input string) bool {
	var status bool
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

// Check if the banner file is valid
func ValidFile(file string) bool {
	fileNames := []string{"standard.txt", "shadow.txt", "thinkertoy.txt"}
	var status bool

	for _, flName := range fileNames {
		if file == flName {
			status = true
		}
	}
	return status
}
