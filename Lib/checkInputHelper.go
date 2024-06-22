package Lib

// Function to trim whitespaces at extreme ends
func trimSpace(s string) string {
	if s == "" {
		return s
	}
	for i := 0; i < len(s); i++ {
		if len(s) > 1 {
			if i == 0 && s[i] == ' ' && s[i+1] != ' ' {
				s = s[1:]
				s = trimSpace(s)
			} else if i == len(s)-1 && s[i] == ' ' && s[i-1] != ' ' {
				s = s[:len(s)-1]
				s = trimSpace(s)
			}
		}
	}
	return s
}

// Detect whether a string starts with a particular set of characters
func hasPrefix(s string, prefix string) bool {
	i := len(prefix)
	return len(s) >= i && s[:i] == prefix
}

// Check if the color flag is valid
func isColorFlag(s string) bool {
	return len(s) > 8 && hasPrefix(s, "--color=")
}

func isFlag(s string) bool {
	return hasPrefix(s, "-") || hasPrefix(s, "--")
}

// Check if flag format is valid
func hasValidFlag(arr []string) {
	for _, s := range arr {
		if isFlag(s) {
			if !isColorFlag(s) {
				PrintError()
			}
		}
	}
}

// Check if flag format is valid
// func markFlag(arr []string) string {
// 	for _, s := range arr {
// 		if isFlag(s) {
// 			return "f"
// 		}
// 	}
// 	return "s"
// }

func isBanner(s string) bool {
	if s == "standard" || s == "shadow" || s == "thinkertoy" {
		return true
	}
	return false
}
