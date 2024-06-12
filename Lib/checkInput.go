package Lib

func CheckInput(input []string) (string, string, string, string, string) {
	// Initialize utility variables
	colorFlag := ""
	subString := ""
	mainString := ""
	color := ""
	reset := ""
	bannerFile := ""

	// Trim spaces from input
	input = inputTrimSpace(input, trimSpace)

	// Check the number of input arguments
	if len(input) == 1 {

		// Exit program if single input is empty
		if len(input[0]) == 0 {
			PrintError()
		}

		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]

			// if argument is has no color flag it is used as main string
		} else {
			mainString = input[0]
			subString = input[0]
		}

		color, reset = colorPicker(colorFlag)

	} else if len(input) == 2 {
		// Exit program if any of the two input strings are empty
		if len(input[0]) == 0 || len(input[1]) == 0 {
			PrintError()
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]
			mainString = input[1]
			subString = input[1]

			// if first argument is a valid banner file assign it as file name, add .txt extension
		} else {
			PrintError()
		}
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 3 {
		// Exit program if any of the three arguments are empty
		if len(input[2]) == 0 || len(input[1]) == 0 || len(input[0]) == 0 {
			PrintError()
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]

			// If third argument is a valid banner file assign it as fileName, add .txt extension
			if input[2] == "standard" || input[2] == "shadow" || input[2] == "thinkertoy" {
				mainString = input[1]
				subString = input[1]
				bannerFile = input[2] + ".txt"
				// Check if user wants to use standard, shadow and thinkertoy as main or sub string
			} else if input[2] == "\\standard" || input[2] == "\\shadow" || input[2] == "\\thinkertoy" {
				mainString = input[2][1:]
				subString = input[1]
			} else {
				mainString = input[2]
				subString = input[1]
			}

			// If first argument is a valid banner file assign it as file name, add .txt extension
		} else {
			PrintError()
		}
		color, reset = colorPicker(colorFlag)

	} else if len(input) == 4 {
		// Exit program if any of the three arguments are empty
		if len(input[0]) == 0 || len(input[1]) == 0 || len(input[2]) == 0 || len(input[3]) == 0 {
			PrintError()
		}

		// If first argument is a valid color flag, assign it to colorFlag
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]

			// If second argument is a valid banner file assign it as fileName, add .txt extension
			if input[3] == "standard" || input[3] == "shadow" || input[3] == "thinkertoy" {
				mainString = input[2]
				subString = input[1]
				bannerFile = input[3] + ".txt"
			} else {
				// Any other format is invalid; print error message, exit program
				PrintError()
			}
		}

		// Extract color and reset codes from colorFlag
		color, reset = colorPicker(colorFlag)
	}
	return color, reset, mainString, subString, bannerFile
}

