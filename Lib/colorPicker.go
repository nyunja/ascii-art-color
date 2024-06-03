package Lib

import (
	"fmt"
	"os"
)

// colorPicker picks the ANSI color code and reset code based on the provided colorFlag.
// It returns the ANSI color code and reset code as strings.
func colorPicker(colorFlag string) (string, string) {
	// Initialize utility variables
	color := ""
	reset := "\033[0m"

	// Define ANSI color codes
	colors := map[string]string{
		"red":     "\033[31m",
		"black":   "\033[30m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"magenta": "\033[35m",
		"blue":    "\033[34m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}

	// Set default color if colorFlag is empty
	if colorFlag == "" {
		color = colors["white"]
	}

	// Check if colorFlag is provided and has a valid format
	if colorFlag != "" && len(colorFlag) < 8 {
		fmt.Println("Error: " + colorFlag + " is not a valid color flag format.")
		os.Exit(0)
	}

	// Handle cases where colorFlag is provided with valid format
	if len(colorFlag) > 8 {
		if hasPrefix(colorFlag, "--color=") {

			key := colorFlag[8:]
			if key == "" {
				color = colors["white"]
			} else {

				// Get ANSI color code based on the provided key
				ansiColorCode, found := colors[key]
				if !found {
					fmt.Println("Error: " + key + " is not a valid color.")
					os.Exit(0)
				}
				color = ansiColorCode
			}
		}
	}

	// Return the ANSI color code and reset code
	return color, reset
}
