package Lib

func toLower(s string) string {
	res := ""
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			res += string(ch+32)
		} else {
			res += string(ch)
		}
	}
	return res
}