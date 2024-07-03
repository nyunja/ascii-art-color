package Lib

type Chars struct {
	index     int
	character rune
}

type Output struct {
	newLine string
	word    []Chars
}

// Returns range of substring to be colored and true if a match is found
// func subToColor(j int, s string, subStrings []string) (bool, int, int) {
// 	for _, sub := range subStrings {
// 		if j <= len(s)-len(sub) && s[j:j+len(sub)] == sub {
// 			return true, j, j + len(sub) - 1
// 		}
// 	}

// 	return false, 0, 0
// }

func getOutputSlice(s string) []Output {
	var outputStructSlice []Output
	var outputStruct Output
	var word []Chars
	var chars Chars
	var countChars int
	for i, ch := range s {
		if ch == '\n' {
			if countChars == len(s[i-countChars:i]) {
				outputStruct.word = word
				outputStructSlice = append(outputStructSlice, outputStruct)
				outputStruct = Output{}
				word = []Chars{}
				countChars = 0
			}
			outputStruct.newLine = "\n"
			outputStructSlice = append(outputStructSlice, outputStruct)
			outputStruct = Output{}
		} else {
			chars.index = i
			chars.character = ch
			word = append(word, chars)
			countChars++
		}
	}
	// fmt.Println("count: ", countChars)
	// fmt.Println("word: ", word)
	outputStruct.word = word
	outputStructSlice = append(outputStructSlice, outputStruct)
	// fmt.Println("output: ", outputStructSlice)
	return outputStructSlice
}

func indicesToColor(s1, s2 string) map[int]bool {
	indices := make(map[int]bool)
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if i+len(s2) <= len(s1) && s1[i:i+len(s2)] == s2 {
				indices = getIndicesToColor(i, i+len(s2), indices)
				i += len(s2) - 1
			}
		}
	}
	return indices
}

func getIndicesToColor(i, j int, indices map[int]bool) map[int]bool {
	for i < j {
		indices[i] = true
		i++
	}
	return indices
}
