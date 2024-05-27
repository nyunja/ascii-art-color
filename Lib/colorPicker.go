package Lib

import (
	"fmt"
	"os"
)

func colorPicker(colorFlag string) (string, string) {
	color := ""
	reset := "\033[0m"

	colors := map[string]string{
		"red":     "\033[31m",
		"black":   "\033[30m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"magenta": "\033[35m",
		"blue":    "\033[34m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	if colorFlag == "" {
		color = colors["white"]
	}
	if colorFlag != "" && len(colorFlag) < 8 {
		fmt.Println("Error: " + colorFlag + " is not a valid color flag format.")
		os.Exit(0)
	}
	if len(colorFlag) > 8 {
		if hasPrefix(colorFlag, "--color=") {

			key := colorFlag[8:]
			if key == "" {
				color = colors["white"]
			} else {
				ansiColorCode, found := colors[key]
				if !found {
					fmt.Println("Error: " + key + " is not a valid color.")
					os.Exit(0)
				}
				color = ansiColorCode
			}
		}
	}

	// else {
	// 	fmt.Println("Usage: go run . [OPTION] [STRING]")
	// 	fmt.Println(`EX: go run . --color=<color> <letters to be colored> "something"`)
	// 	os.Exit(0)
	// }
	return color, reset
}
