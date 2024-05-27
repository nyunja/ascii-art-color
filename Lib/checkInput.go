package Lib

func CheckInput(input []string) (string, string, string, string) {
	colorFlag := ""
	subString := ""
	mainString := ""
	color := ""
	reset := ""
	if len(input) == 1 {
		mainString = input[0]
		subString = input[0]
		color, reset = colorPicker(colorFlag)
		// fmt.Printf("main : %s sun : %s  color : %s reset : %s", mainString, subString, color, reset)
		// printOut(color, reset, mainString, subString)
	} else if len(input) == 2 {
		colorFlag = input[0]
		mainString = input[1]
		subString = input[1]
		color, reset = colorPicker(colorFlag)
		// printOut(color, reset, mainString, subString)
	} else if len(input) == 3 {
		colorFlag = input[0]
		subString = input[1]
		mainString = input[2]
		// for _, ch := range mainString {
		// 	if !matchingStr(subString, string(ch)) {
		// 		fmt.Printf("Error: %s has no match\n", subString)
		// 		os.Exit(0)
		// 	}
		// }
		color, reset = colorPicker(colorFlag)
		// printOut(color, reset, mainString, subString)
	}
	return color, reset, mainString, subString
}

// func printOut(color string, reset string, mainString string, subString string) {
// 	match, found := matchingStr(mainString, subString)
// 	if !found {
// 		fmt.Printf("Error: %s has no match\n", subString)
// 		os.Exit(0)
// 	}
// 	// for _, ch := range mainString {
// 	// 	if matchingStr(subString, string(ch)) && !match[string(ch)] {
// 	// 		match[string(ch)] = true
// 	// 	} else if !matchingStr(subString, string(ch)) {
// 	// 		match[string(ch)] = false
// 	// 	}
// 	// 	// else if !matchingStr(subString, string(ch)) && i == len(mainString) -1 {
// 	// 	// 	fmt.Printf("Error: %s has no match\n", subString)
// 	// 	// 	os.Exit(0)
// 	// 	// }
// 	// }
// 	for _, ch := range mainString {
// 		if match[string(ch)] {
// 			fmt.Printf("%s%s%s", color, string(ch), reset)
// 		} else {
// 			fmt.Printf("%s", string(ch))
// 			// else if matchingStr(subString, string(ch)) && match[string(ch)] {
// 			// 	fmt.Printf("%s%s%s", color, string(ch), reset)
// 			// } else if !matchingStr(subString, string(ch)) && !match[string(ch)] {
// 			// 	fmt.Printf("%s", string(ch))
// 			// }
// 		}
// 	}
// 	fmt.Println()
// }

func matchingStr(s string, subString string) (map[string]bool, bool) {
	match := make(map[string]bool)
	for _, ch := range s {
		for _, ch1 := range subString {
			if ch1 == ch {
				match[string(ch1)] = true
			}
		}
	}

	if len(match) > 0 {
		return match, true
	}
	return nil, false
}

// func hasSuffix(s string, suffix string) bool {
// 	i := len(s) - len(suffix)
// 	return s[i:] == suffix
// }

func hasPrefix(s string, prefix string) bool {
	i := len(prefix)
	return s[:i] == prefix
}
