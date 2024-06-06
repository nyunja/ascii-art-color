package Lib

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
		"orange":  "\033[38;5;214m",
		"indigo":  "\033[38;5;57m",
		"violet":  "\033[38;5;129m",
		"purple":  "\033[38;5;129m",
		"pink":    "\033[38;5;218m",
		"brown":   "\033[38;5;130m",
		"olive":   "\033[38;5;58m",
		"teal":    "\033[38;5;30m",
		"maroon":  "\033[38;5;52m",
	}

	// Set default color if colorFlag is empty
	if colorFlag == "" {
		color = colors["white"]
	}

	// Check if colorFlag is provided and has a valid format
	if colorFlag != "" && len(colorFlag) < 8 {
		PrintError()
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
					PrintError()
				}
				color = ansiColorCode
			}
		}
	}

	// Return the ANSI color code and reset code
	return color, reset
}
