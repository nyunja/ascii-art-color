package Lib

import (
	"fmt"
	"os"
	"strings"
)

func getFinalOutput(color1, color2, reset, mainStr, subStr string, lines []string) string {
	var output string
	var startIdx, charCount, countNewlines int

	//outputStructSlice := getOutputSlice(mainStr)      // Extract characters to be printed
	indicesToColor := indicesToColor(mainStr, subStr) // Extract indices that need coloring
	fmt.Println(indicesToColor)

	// Populate output with ascii-art
	words := strings.Split(mainStr, "\n")
	fmt.Printf("%q\n", words)
	for _, word := range words {
		if word == "" {
			countNewlines++
			if countNewlines < len(words) {
				output += "\n"
				charCount++
				fmt.Println("newline: ", charCount)
			}
		} else {
			//word := word.word
			if len(word) > 0 {
				// Populate output with every line of ascii-art
				for k := 0; k < 8; k++ {
					for _, ch := range word {
						startIdx = int(ch-32)*9 + 1

						// Append color code for color1 to appropriate characters
						if len(color1) > 0 && indicesToColor[charCount] {
							fmt.Println("true: ", charCount)
							output += color1 + lines[startIdx+k] + reset
							charCount++

							// if true, execute color code of color2
						} else if len(color2) > 0 {
							output += color2 + lines[startIdx+k] + reset
							charCount++

							// if no index to color, append ascii art without color codes
						} else {
							output += lines[startIdx+k]
							charCount++
						}
					}
					charCount = charCount - len(word)
					fmt.Println(charCount)

					// Add new-line-character at the end of each line of output
					output += "\n"
				}
				charCount = charCount + len(word)
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
