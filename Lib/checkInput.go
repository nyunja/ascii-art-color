package Lib

import (
	"fmt"
	"os"
)

func CheckInput(input []string) (string, string, string, string, string) {
	// Initialize utility variables
	colorFlag := ""
	subString := ""
	mainString := ""
	color := ""
	reset := ""
	fileName := ""

	// Check the number of input arguments
	if len(input) == 1 {

		// Exit program if single input is empty
		if len(input[0]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}
		mainString = input[0]
		subString = input[0]
		color, reset = colorPicker(colorFlag)

	} else if len(input) == 2 {
		// Exit program if any of the two input strings are empty
		if len(input[0]) == 0 || len(input[1]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]
			mainString = input[1]
			subString = input[1]

			// if first argument is a valid banner file assign it as file name, add .txt extension
		} else if input[0] == "-standard" || input[0] == "-shadow" || input[0] == "-thinkertoy" {
			mainString = input[1]
			subString = input[1]
			fileName = input[0][1:] + ".txt"

			// Any other format is invalid; print error message, exit program
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
			os.Exit(0)
		}
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 3 {
		// Exit program if any of the three arguments are empty
		if len(input[2]) == 0 || len(input[1]) == 0 || len(input[0]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]

			// If second argument is a valid banner file assign it as fileName, add .txt extension
			if input[1] == "-standard" || input[1] == "-shadow" || input[1] == "-thinkertoy" {
				mainString = input[2]
				subString = input[2]
				fileName = input[1][1:] + ".txt"
			} else {

				mainString = input[2]
				subString = input[1]
			}

			// If first argument is a valid banner file assign it as file name, add .txt extension
		} else if input[0] == "-standard" || input[0] == "-shadow" || input[0] == "-thinkertoy" {
			mainString = input[2]
			subString = input[1]
			fileName = input[0][1:] + ".txt"

			// Any other format is invalid; print error message, exit program
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
			os.Exit(0)
		}
		color, reset = colorPicker(colorFlag)

	} else if len(input) == 4 {
		// Exit program if any of the three arguments are empty
		if len(input[0]) == 0 || len(input[1]) == 0 || len(input[2]) == 0 || len(input[3]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]

			// If second argument is a valid banner file assign it as fileName, add .txt extension
			if input[1] == "-standard" || input[1] == "-shadow" || input[1] == "-thinkertoy" {
				mainString = input[3]
				subString = input[2]
				fileName = input[1][1:] + ".txt"
			} else {
				// Any other format is invalid; print error message, exit program
				fmt.Println("Usage: go run . [OPTION] [STRING]")
				fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
				os.Exit(0)
			}
		}

		// Extract color and reset codes from colorFlag
		color, reset = colorPicker(colorFlag)
	}

	return color, reset, mainString, subString, fileName
}

// matchingStr checks if any character in the string s matches any character in the substring subString.
// It returns a map indicating which characters match, and a boolean indicating if there was any match.
func matchingStr(s string, subString string) (map[string]bool, bool) {

	// Initialize a map to store matches
	match := make(map[string]bool)

	// Loop through each character in the string s, and subsequently through subString
	for _, ch := range s {
		for _, ch1 := range subString {

			// If a character in subString matches a character in s, mark it as a match
			if ch1 == ch {
				match[string(ch1)] = true
			}
		}
	}

	// If there are matches, return the map of matches and true
	if len(match) > 0 {
		return match, true
	}

	// If there are no matches, return nil map and false
	return nil, false
}

// Detect whether a string starts with a particular set of characters
func hasPrefix(s string, prefix string) bool {
	i := len(prefix)
	return s[:i] == prefix
}
