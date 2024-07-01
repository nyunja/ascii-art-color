package Lib

// Function to trim whitespaces at extreme ends
func trimSpace(s string) string {
	if s == "" {
		return s
	}
	for i := 0; i < len(s); i++ {
		if len(s) > 1 {
			// Remove space if it appears as the first character of a string
			if i == 0 && s[i] == ' ' && s[i+1] != ' ' {
				s = s[1:]
				s = trimSpace(s)

				// Remove space if it appears as the last character of a string
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

// Detect whether a string starts with a particular set of characters
func hasSuffix(s string, suffix string) bool {
	i := len(s) - len(suffix)
	return i >= 0 && s[i:] == suffix
}

// Check if the color flag is valid
func isColorFlag(s string) bool {
	return len(s) > 8 && hasPrefix(s, "--color=")
}

// Check if the output flag is valid
func isOutputFlag(s string) bool {
	return len(s) > 9 && hasPrefix(s, "--output=") && hasSuffix(s, ".txt")
}

// Identifis strings that start with '-' or '--'
func isFlag(s string) bool {
	return hasPrefix(s, "-") || hasPrefix(s, "--")
}

// Identifies whether string is a banner signal
func isBanner(s string) bool {
	if s == "standard" || s == "shadow" || s == "thinkertoy" {
		return true
	}
	return false
}

// Modify the main string by adding characters or replacing newline sequences
func unmatchSubstring(s string) string {
	if containsString(s, "\\n") {
		s = replaceString(s, "\\n", "\\ns")
	} else {
		s = "s" + s + "s"
	}
	return s
}

// Check if a substring exists within a string
func containsString(s, sub string) bool {
	for i := 0; i < len(s); i++ {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// Replace occurrences of a substring with another substring
func replaceString(s, sub, rep string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			res += rep
			i += len(sub) - 1
		}
		res += string(s[i])
	}
	return res
}
