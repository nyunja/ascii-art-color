package Lib

// type flagz struct {
// 	s string
// 	t string
// }

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
	// fmt.Println(input)

	// inputs := []flagz{}
	// inputs := input

	// for _, str := range input {
	// 	if isFlag(str) {
	// 		flg = append(flg, flagz{str, "f"})
	// 	} else {
	// 		flg = append(flg, flagz{str, "s"})
	// 	}
	// }
	// fmt.Println(flg)
	// colorFlag, outputFile, bannerFile, flg = sortInput(colorFlag, outputFile, bannerFile, flg)
	colorFlag, outputFile, bannerFile, inputs = sortInput2(colorFlag, outputFile, bannerFile, inputs)

	if len(inputs) < 1 || len(inputs) > 2 {
		PrintError()
	}
	if len(inputs) == 2 {
		if inputs[1] == "\\standard" || inputs[1] == "\\shadow" || inputs[1] == "\\thinkertoy" {
			inputs[1] = inputs[1][1:]
		}
		mainString = inputs[1]
		subString = inputs[0]
	} else if len(inputs) == 1 {
		mainString = inputs[0]
		subString = inputs[0]
	}

	// if len(flg) == 2 {
	// 	if flg[1].s == "\\standard" || flg[1].s == "\\shadow" || flg[1].s == "\\thinkertoy" {
	// 		flg[1].s = flg[1].s[1:]
	// 	}
	// 	mainString = flg[1].s
	// 	subString = flg[0].s
	// } else if len(flg) == 1 {
	// 	mainString = flg[0].s
	// 	subString = flg[0].s
	// }

	// // Check if flags are valid
	// hasValidFlag(input)

	// // Check the number of input arguments
	// if len(input) == 1 {
	// 	if isColorFlag(input[0]) {
	// 		PrintError()
	// 	}
	// 	// If argument is has no color flag it is used as main string
	// 	mainString = input[0]
	// 	subString = input[0]
	// } else if len(input) == 2 {
	// 	// If first argument is a valid color flag, assign it to colorFlag
	// 	if isColorFlag(input[0]) {
	// 		colorFlag = input[0]
	// 		mainString = input[1]
	// 		subString = input[1]
	// 		// If second argument is a valid banner file assign it as fileName, add .txt extension
	// 	} else if isBanner(input[1]) {
	// 		mainString = input[0]
	// 		subString = input[0]
	// 		bannerFile = input[1] + ".txt"
	// 		// if first argument is a valid banner file assign it as file name, add .txt extension
	// 	} else {
	// 		PrintError()
	// 	}
	// } else if len(input) == 3 {
	// 	// If first argument is a valid color flag, assign it to colorFlag
	// 	if isColorFlag(input[0]) {
	// 		colorFlag = input[0]
	// 		// If third argument is a valid banner file assign it as fileName, add .txt extension
	// 		if isBanner(input[2]) {
	// 			mainString = input[1]
	// 			subString = input[1]
	// 			bannerFile = input[2] + ".txt"
	// 			// Check if user wants to use standard, shadow and thinkertoy as main or sub string
	// 		} else if input[2] == "\\standard" || input[2] == "\\shadow" || input[2] == "\\thinkertoy" {
	// 			mainString = input[2][1:]
	// 			subString = input[1]
	// 		} else {
	// 			mainString = input[2]
	// 			subString = input[1]
	// 		}
	// 		// If first argument is a valid banner file assign it as file name, add .txt extension
	// 	} else {
	// 		PrintError()
	// 	}
	// } else if len(input) == 4 {
	// 	// If first argument is a valid color flag, assign it to colorFlag
	// 	if isColorFlag(input[0]) {
	// 		colorFlag = input[0]
	// 		// If second argument is a valid banner file assign it as fileName, add .txt extension
	// 		if isBanner(input[3]) {
	// 			mainString = input[2]
	// 			subString = input[1]
	// 			bannerFile = input[3] + ".txt"
	// 		}
	// 	} else {
	// 		PrintError()
	// 	}
	// }
	// Extract color and reset codes from colorFlag
	color1, color2, reset = colorPicker(colorFlag)

	return color1, color2, reset, mainString, subString, bannerFile, outputFile
}

// func sortInput(colorFlag string, outputFile string, bannerFile string, flg []flagz) (string, string, string, []flagz) {
// 	for i, item := range flg {
// 		if item.t == "f" {
// 			if i < len(flg)-1 && isColorFlag(item.s) {
// 				colorFlag = item.s
// 				flg = append(flg[:i], flg[i+1:]...)
// 				colorFlag, outputFile, bannerFile, flg = sortInput(colorFlag, outputFile, bannerFile, flg)
// 			} else if i < len(flg)-1 && isOutputFlag(item.s) {
// 				outputFile = item.s[9:]
// 				flg = append(flg[:i], flg[i+1:]...)
// 				colorFlag, outputFile, bannerFile, flg = sortInput(colorFlag, outputFile, bannerFile, flg)
// 			} else {
// 				PrintError()
// 			}
// 		} else if i == len(flg)-1 && item.t == "s" {
// 			if isBanner(item.s) {
// 				bannerFile = item.s + ".txt"
// 				flg = flg[:i]
// 			}
// 		}
// 	}
// 	return colorFlag, outputFile, bannerFile, flg
// }

func sortInput2(colorFlag string, outputFile string, bannerFile string, flg []string) (string, string, string, []string) {
	for i, item := range flg {
		if isFlag(item) {
			if i < len(flg)-1 && isColorFlag(item) {
				colorFlag = item
				flg = append(flg[:i], flg[i+1:]...)
				colorFlag, outputFile, bannerFile, flg = sortInput2(colorFlag, outputFile, bannerFile, flg)
			} else if i < len(flg)-1 && isOutputFlag(item) {
				outputFile = item[9:]
				flg = append(flg[:i], flg[i+1:]...)
				colorFlag, outputFile, bannerFile, flg = sortInput2(colorFlag, outputFile, bannerFile, flg)
			} else {
				PrintError()
			}
		} else if i == len(flg)-1 && isBanner(item) {
			bannerFile = item + ".txt"
			flg = flg[:i]
		}
	}
	return colorFlag, outputFile, bannerFile, flg
}
