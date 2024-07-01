package Lib

// Conert the first character of string parsed to lowercase
func toLower(s string) string {
	res := ""
	for _, ch := range s {
		// match capital letters and convert them to lowercase
		if ch >= 'A' && ch <= 'Z' {
			res += string(ch + 32)
		} else {
			res += string(ch)
		}
	}
	return res
}
