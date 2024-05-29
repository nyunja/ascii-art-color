package Lib

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(color, reset, mainString, subString, fileName string) string {
	// Exits program if mainString or subString is empty
	if mainString == "" || subString == "" {
		return ""
	}

	// Exits program if input contains characters outside of the banner file (except '\n')
	if !IsPrintable(mainString) || !IsPrintable(subString) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}

	// Exits program if input contains non-printable ascii characters with escape sequences (except '\n')
	if EscapeSequence(mainString) || EscapeSequence(subString) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}
	// Make new line characters consistent
	mainString = strings.ReplaceAll(mainString, "\\n", "\n")

	// Split input into printable lines at the '\n' character
	words := strings.Split(mainString, "\n")

	// Set and check if banner file is valid
	file := ""
	if len(fileName) == 0 {
		file = "standard.txt"
	} else {
		file = fileName
	}
	if !ValidFile(file) {
		fmt.Println("Incorrect file name, program only accepts thinkertoy.txt, standard.txt or shadow.txt")
		return ""
	}

	// Read banner file
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading %s\n", file)
		return ""
	}

	// Split and store the file contents line-by-line in a slice string depending on the banner file
	var slices []string
	if file == "thinkertoy.txt" {
		// Lines in thinkertoy.txt banner file are separated by "\r\n"
		slices = strings.Split(string(content), "\r\n")
	} else {
		// Lines in standard.txt and shadow.txt banner files are separated by "\n"
		slices = strings.Split(string(content), "\n")
	}

	// Exit program if banner file is incomplete
	if len(slices) != 856 {
		fmt.Printf("Invalid file content or content modified, check %s\n", file)
		return ""
	}

	// Print ASCII ART and return output string for testing
	return HandleWords(color, reset, subString, slices, words)
}
