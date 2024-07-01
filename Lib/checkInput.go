package Lib

func CheckInput(inputs []string) (string, string, string, string, string, string, string) {
	// Initialize utility variables
	colorFlag := ""
	subString := ""
	mainString := ""
	color1 := ""
	color2 := ""
	reset := ""
	bannerFile := ""
	outputFile := ""

	// Trim spaces from input
	inputs = inputTrimSpace(inputs, trimSpace)
	colorFlag, outputFile, bannerFile, inputs = sortInput(colorFlag, outputFile, bannerFile, inputs)

	// Print error message when extra input strings are detected
	if len(inputs) < 1 || len(inputs) > 2 {
		PrintError()
	}

	if len(inputs) == 2 {
		// Check if user input sub and mainstring without color or other flags, print error massage
		if len(bannerFile) == 0 && len(colorFlag) == 0 && len(outputFile) == 0 {
			PrintError()
		}

		// Shave-off the escape character when banner signals are used a the mainstring
		if inputs[1] == "\\standard" || inputs[1] == "\\shadow" || inputs[1] == "\\thinkertoy" {
			inputs[1] = inputs[1][1:]
		}
		mainString = inputs[1]
		subString = inputs[0]

		//When user parses empty sub-string add extra character to unmatch it to mainstring
		if len(subString) == 0 {
			subString = unmatchSubstring(mainString)
		}

	} else if len(inputs) == 1 {
		mainString = inputs[0]
		subString = inputs[0]
	}

	// Extract color and reset codes from colorFlag
	color1, color2, reset = colorPicker(colorFlag)

	return color1, color2, reset, mainString, subString, bannerFile, outputFile
}

// Sort through arguments, labeling appropriate flags and strings and sub-string
func sortInput(colorFlag string, outputFile string, bannerFile string, input []string) (string, string, string, []string) {
	for i, item := range input {
		if isFlag(item) {
			// Detect and label color flag
			if i < len(input)-1 && isColorFlag(item) {
				colorFlag = item
				input = append(input[:i], input[i+1:]...)
				colorFlag, outputFile, bannerFile, input = sortInput(colorFlag, outputFile, bannerFile, input)

				// Detect and label output flag
			} else if i < len(input)-1 && isOutputFlag(item) {
				outputFile = item[9:]
				input = append(input[:i], input[i+1:]...)
				colorFlag, outputFile, bannerFile, input = sortInput(colorFlag, outputFile, bannerFile, input)

				//Anything must have an invalid flag format, print error
			} else {
				PrintError()
			}

			// Add '.txt' extension to banner input
		} else if len(bannerFile) == 0 && i == len(input)-1 && isBanner(item) {
			bannerFile = item + ".txt"
			input = input[:i]
		}
	}
	return colorFlag, outputFile, bannerFile, input
}
