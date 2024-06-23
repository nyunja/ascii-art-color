package main

import (
	"fmt"
	"os"

	"colors/Lib"
)

func main() {
	// Saves the input from the user
	inputs := os.Args[1:]

	// Exits the program if the arguments passed are greater than 5
	if len(inputs) > 5 {
		Lib.PrintError()
	}
	// Exits the program if the arguments passed are less than 1
	if len(inputs) == 0 {
		Lib.PrintError()
	}

	// Check what user input contains and returns required variables
	color1, color2, reset, mainString, subString, fileName, outputFile := Lib.CheckInput(inputs)

	// Call the AsciiArt function to handle input
	output := Lib.AsciiArt(color1, color2, reset, mainString, subString, fileName)

	if len(outputFile) > 0 {
		os.WriteFile(outputFile, []byte(output), 0666)
	} else {
		fmt.Print(output)
	}
}
