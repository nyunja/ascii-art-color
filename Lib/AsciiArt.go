package Lib

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(color1, color2, reset, mainString, subString, bannerFile string) string {
	// Exits program if mainString or subString is empty
	if mainString == "" || subString == "" {
		return ""
	}

	// Exits program if input contains characters outside of the banner file (except '\n')
	if !IsPrintable(mainString) || !IsPrintable(subString) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}

	// Exits program if input contains non-printable ascii characters with escape sequences (except '\n')
	if EscapeSequence(mainString) || EscapeSequence(subString) {
		fmt.Println(`Input should only contain PRINTABLE ASCII characters or '\n'`)
		return ""
	}
	// Make new line characters consistent
	mainString = strings.ReplaceAll(mainString, "\\n", "\n")
	subString = strings.ReplaceAll(subString, "\\n", "\n")

	// Split inputs into printable lines at the '\n' character
	words := strings.Split(mainString, "\n")
	subStrings := strings.Split(subString, "\n")

	// Set and check if banner file is valid
	file := ""
	if len(bannerFile) == 0 {
		file = "standard.txt"
	} else {
		file = bannerFile
	}

	// Read banner file
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading %s\n", file)
		return ""
	}

	// Compare checksums
	h := sha256.New()
	h.Write(content)
	s := fmt.Sprintf("%x", h.Sum(nil))

	shadowChecksum := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	standardChecksum := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	toyChecksum := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"

	// Compare checksums
	if s != shadowChecksum && s != standardChecksum && s != toyChecksum {
		fmt.Printf("Invalid file content or content modified, check %s\n", file)
		return ""
	}

	// Split and store the file contents line-by-line in a slice string depending on the banner file
	var slices []string
	if file == "thinkertoy.txt" {
		// Lines in thinkertoy.txt banner file are separated by "\r\n"
		slices = strings.Split(string(content), "\r\n")
	} else {
		// Lines in standard.txt and shadow.txt banner files are separated by "\n"
		slices = strings.Split(string(content), "\n")
	}

	// Exit program if banner file is incomplete
	if len(slices) != 856 {
		fmt.Printf("Invalid file content or content modified, check %s\n", file)
		return ""
	}

	// Print ASCII ART and return output string for testing
	return HandleWords(color1, color2, reset, subStrings, slices, words)
}
