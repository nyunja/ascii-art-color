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

	if len(inputs) < 1 || len(inputs) > 2 {
		PrintError()
	}
	if len(inputs) == 2 {
		if len(bannerFile) == 0 && len(colorFlag) == 0 && len(outputFile) == 0 {
			PrintError()
		}
		if inputs[1] == "\\standard" || inputs[1] == "\\shadow" || inputs[1] == "\\thinkertoy" {
			inputs[1] = inputs[1][1:]
		}
		mainString = inputs[1]
		subString = inputs[0]
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

func sortInput(colorFlag string, outputFile string, bannerFile string, flg []string) (string, string, string, []string) {
	for i, item := range flg {
		if isFlag(item) {
			if i < len(flg)-1 && isColorFlag(item) {
				colorFlag = item
				flg = append(flg[:i], flg[i+1:]...)
				colorFlag, outputFile, bannerFile, flg = sortInput(colorFlag, outputFile, bannerFile, flg)
			} else if i < len(flg)-1 && isOutputFlag(item) {
				outputFile = item[9:]
				flg = append(flg[:i], flg[i+1:]...)
				colorFlag, outputFile, bannerFile, flg = sortInput(colorFlag, outputFile, bannerFile, flg)
			} else {
				PrintError()
			}
		} else if len(bannerFile) == 0 && i == len(flg)-1 && isBanner(item) {
			bannerFile = item + ".txt"
			flg = flg[:i]
		}
	}
	return colorFlag, outputFile, bannerFile, flg
}
