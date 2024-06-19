package main

import (
	"fmt"
	"os"
	"strings"

	"colors/Lib"
)

func main() {
	// Saves the input from the user
	input := os.Args[1:]

	// Exits the program if the arguments passed are greater than 5
	if len(input) > 4 {
		Lib.PrintError()
	}
	// Exits the program if the arguments passed are less than 1
	if len(input) == 0 {
		Lib.PrintError()
	}

	// // Check what user input contains and returns required variables
	color, reset, mainString, subString, fileName := Lib.CheckInput(input)
	// fmt.Printf("main : %s sun : %s  color : %s reset : %s", mainString, subString, color, reset)

	// Call the AsciiArt function to handle input
	lines := Lib.AsciiArt(color, reset, mainString, subString, fileName)

	for _, line := range strings.Split(lines, "\n") {
		fmt.Printf("%q\n", line)
	}
}
