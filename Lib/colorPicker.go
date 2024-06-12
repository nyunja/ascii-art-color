package Lib

// colorPicker picks the ANSI color code and reset code based on the provided colorFlag.
// It returns the ANSI color code and reset code as strings.
func colorPicker(colorFlag string) (string, string) {
	// Initialize utility variables
	color := ""
	reset := "\033[0m"

	// Define ANSI color codes
	colors := map[string]string{
		"red":       "\033[31m",
		"green":     "\033[32m",
		"blue":      "\033[34m",
		"orange":    "\033[38;5;208m",
		"yellow":    "\033[33m",
		"black":     "\033[30m",
		"white":     "\033[37m",
		"gray":      "\033[90m",
		"pink":      "\033[95m",
		"teal":      "\033[36m",
		"purple":    "\033[35m",
		"brown":     "\033[33;2m",
		"beige":     "\033[33;2m",
		"tan":       "\033[33;2m",
		"peach":     "\033[95m",
		"lime":      "\033[92m",
		"olive":     "\033[92m",
		"turquoise": "\033[96m",
		"navyBlue":  "\033[34;1m",
		"indigo":    "\033[94m",
		"violet":    "\033[94m",
		"lavender":  "\033[94m",
		"lilac":     "\033[94m",
		"maroon":    "\033[31;2m",
		"crimson":   "\033[31;2m",
		"fuchsia":   "\033[95m",
		"magenta":   "\033[95m",
		"coral":     "\033[95m",
		"saffron":   "\033[93m",
		"sage":      "\033[92m",
		"mint":      "\033[92m",
		"seafoam":   "\033[96m",
		"emerald":   "\033[92m",
		"sapphire":  "\033[94m",
		"ruby":      "\033[91m",
		"garnet":    "\033[91m",
		"ivory":     "\033[97m",
		"cream":     "\033[97m",
		"champagne": "\033[97m",
		"camel":     "\033[33;2m",
		"khaki":     "\033[33;2m",
		"mocha":     "\033[33;2m",
		"umber":     "\033[33;2m",
		"sienna":    "\033[33;2m",
		"charcoal":  "\033[90m",
		"slate":     "\033[90m",
		"ash":       "\033[90m",
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

			key := trimSpace(colorFlag[8:])
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
