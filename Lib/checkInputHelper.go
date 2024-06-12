package Lib

// Function to trim whitespaces at extreme ends
func trimSpace(s string) string {
	if s == "" {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == ' ' {
			s = s[1:]
			s = trimSpace(s)
		} else if i == len(s)-1 && s[i] == ' ' {
			s = s[:len(s)-1]
			s = trimSpace(s)
		}
	}
	return s
}

// Detect whether a string starts with a particular set of characters
func hasPrefix(s string, prefix string) bool {
	i := len(prefix)
	return s[:i] == prefix
}