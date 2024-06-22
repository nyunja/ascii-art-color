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

	// Check if flags are valid
	hasValidFlag(input)

	// Check the number of input arguments
	if len(input) == 1 {
		if isColorFlag(input[0]) {
			PrintError()
		}
		// If argument is has no color flag it is used as main string
		mainString = input[0]
		subString = input[0]
	} else if len(input) == 2 {
		// If first argument is a valid color flag, assign it to colorFlag
		if isColorFlag(input[0]) {
			colorFlag = input[0]
			mainString = input[1]
			subString = input[1]
			// If second argument is a valid banner file assign it as fileName, add .txt extension
		} else if isBanner(input[1]) {
			mainString = input[0]
			subString = input[0]
			bannerFile = input[1] + ".txt"
			// if first argument is a valid banner file assign it as file name, add .txt extension
		} else {
			PrintError()
		}
	} else if len(input) == 3 {
		// If first argument is a valid color flag, assign it to colorFlag
		if isColorFlag(input[0]) {
			colorFlag = input[0]
			// If third argument is a valid banner file assign it as fileName, add .txt extension
			if isBanner(input[2]) {
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
	} else if len(input) == 4 {
		// If first argument is a valid color flag, assign it to colorFlag
		if isColorFlag(input[0]) {
			colorFlag = input[0]
			// If second argument is a valid banner file assign it as fileName, add .txt extension
			if isBanner(input[3]) {
				mainString = input[2]
				subString = input[1]
				bannerFile = input[3] + ".txt"
			}
		} else {
			PrintError()
		}
	}
	// Extract color and reset codes from colorFlag
	color, reset = colorPicker(colorFlag)

	return color, reset, mainString, subString, bannerFile
}
