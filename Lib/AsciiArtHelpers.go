package Lib

import "strings"

// Check if input contains non-printable ascii characters with escape sequences (except '\n')
func EscapeSequence(input string) bool {
	allowedEscapes := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r"}
	for _, unprint := range allowedEscapes {
		// Detect escape sequence characters in input string
		if strings.Contains(input, unprint) {
			return true
		}
	}
	return false
}

// Check if input contains characters outside of the banner file (except '\n')
func IsPrintable(input string) bool {
	var status bool

	// Check if character is represented in banner file
	for _, v := range input {
		if (v >= ' ' && v <= '~') || v == '\n' {
			status = true
		} else {
			status = false
			break
		}
	}
	return status
}

// Trim extra spaces from input strings
func inputTrimSpace(input []string, f func(string) string) []string {
	var res []string
	for _, s := range input {
		res = append(res, f(s))
	}
	return res
}
