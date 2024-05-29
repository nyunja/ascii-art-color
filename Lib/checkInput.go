package Lib

import (
	"fmt"
	"os"
)

func CheckInput(input []string) (string, string, string, string, string) {
	colorFlag := ""
	subString := ""
	mainString := ""
	color := ""
	reset := ""
	fileName := ""
	if len(input) == 1 {
		if len(input[0]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}
		mainString = input[0]
		subString = input[0]
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 2 {
		if len(input[0]) == 0 || len(input[1]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]
			mainString = input[1]
			subString = input[1]
		} else if input[0] == "-standard" || input[0] == "-shadow" || input[0] == "-thinkertoy" {
			mainString = input[1]
			subString = input[1]
			fileName = input[0][1:] + ".txt"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
			os.Exit(0)
		}
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 3 {
		if len(input[2]) == 0 || len(input[1]) == 0 || len(input[0]) == 0  {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]
			if input[1] == "-standard" || input[1] == "-shadow" || input[1] == "-thinkertoy" {
				mainString = input[2]
				subString = input[2]
				fileName = input[1][1:] + ".txt"
			} else {
				mainString = input[2]
				subString = input[1]
			}
		} else if input[0] == "-standard" || input[0] == "-shadow" || input[0] == "-thinkertoy" {
			mainString = input[2]
			subString = input[1]
			fileName = input[0][1:] + ".txt"
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
			os.Exit(0)
		}
		color, reset = colorPicker(colorFlag)
	} else if len(input) == 4 {
		if len(input[0]) == 0 || len(input[1]) == 0 || len(input[2]) == 0 || len(input[3]) == 0 {
			fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
			os.Exit(0)
		}
		if len(input[0]) >= 8 && hasPrefix(input[0], "--color=") {
			colorFlag = input[0]
			if input[1] == "-standard" || input[1] == "-shadow" || input[1] == "-thinkertoy" {
				mainString = input[3]
				subString = input[2]
				fileName = input[1][1:] + ".txt"
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING]")
				fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
				os.Exit(0)
			}
			// else {
			// 	mainString = input[2]
			// 	subString = input[1]
			// }
		}
		// else if input[0] == "-standard" || input[0] == "-shadow" || input[0] == "-thinkertoy" {
		// 	if len(input[1]) == 0 {
		// 		fmt.Println("error: arguments cannot be empty strings. please provide inputs.")
		// 		os.Exit(0)
		// 	}
		// 	mainString = input[2]
		// 	subString = input[1]
		// 	fileName = input[0][1:] + ".txt"
		// } else {
		// 	fmt.Println("Usage: go run . [OPTION] [STRING]")
		// 	fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
		// 	os.Exit(0)
		// }
		// colorFlag = input[0]
		// subString = input[1]
		// mainString = input[2]
		color, reset = colorPicker(colorFlag)
	}
	return color, reset, mainString, subString, fileName
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

func hasPrefix(s string, prefix string) bool {
	i := len(prefix)
	return s[:i] == prefix
}
