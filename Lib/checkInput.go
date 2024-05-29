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
	} else if len(input) == 2 {
		colorFlag = input[0]
		mainString = input[1]
		subString = input[1]
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 3 {
		colorFlag = input[0]
		subString = input[1]
		mainString = input[2]
		color, reset = colorPicker(colorFlag)
	}
	return color, reset, mainString, subString
}

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
