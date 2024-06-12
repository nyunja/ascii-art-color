package Lib

// Returns range of substring to be colored and true if a match is found
func lineToColor(s string, subStrings []string) (bool, int, int) {
	for _, sub := range subStrings {
		for i := 0; i < len(s); i++ {
			if i <= len(s)-len(sub) && s[i:i+len(sub)] == sub {
				return true, i, i + len(sub) - 1
			}
		}
	}

	return false, 0, 0
}

// Returns range of substring to be colored and true if a match is found
func subToColor(j int, s string, subStrings []string) (bool, int, int) {
	for _, sub := range subStrings {
		if j <= len(s)-len(sub) && s[j:j+len(sub)] == sub {
			return true, j, j + len(sub) - 1
		}
	}

	return false, 0, 0
}
