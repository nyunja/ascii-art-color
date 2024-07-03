package Lib

import (
	"fmt"
	"os"
)

func getFinalOutput(color1, color2, reset, s1, s2 string, lines []string) string {
	var output string
	var startIdx int
	outputStructSlice := getOutputSlice(s1)
	indicesToColor := indicesToColor(s1, s2)
	for _, outputStruct := range outputStructSlice {
		if outputStruct.newLine == "\n" {
			output += "\n"
		} else {
			word := outputStruct.word
			if len(word) > 0 {
				for k := 0; k < 8; k++ {
					for _, ch := range word {
						startIdx = int(ch.character-32)*9 + 1
						if len(color1) > 0 && indicesToColor[ch.index] {
							output += color1 + lines[startIdx+k] + reset
						} else if len(color2) > 0 {
							output += color2 + lines[startIdx+k] + reset
						} else {
							output += lines[startIdx+k]
						}
					}
					output += "\n"
				}
			}
		}
	}
	return output
}

// Prints the standard error message
func PrintError() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	os.Exit(0)
}
