package Lib

type Chars struct {
	index     int
	character rune
}

type Output struct {
	newLine string
	word    []Chars
}

func getOutputSlice(s string) []Output {
	var outputStructSlice []Output
	var word []Chars
	var countNewlines int
	for i, ch := range s {
		if ch == '\n' {
			countNewlines++
			if len(word) > 0 {
				outputStructSlice = append(outputStructSlice, Output{word: word})
				word = nil
			}
			if countNewlines > 1 {
				outputStructSlice = append(outputStructSlice, Output{newLine: "\n"})
			}
			if countNewlines == len(s) {
				outputStructSlice = append(outputStructSlice, Output{newLine: "\n"})
				break
			}
		} else {
			word = append(word, Chars{index: i, character: ch})
			countNewlines = 0
		}
	}
	if len(word) > 0 {
		outputStructSlice = append(outputStructSlice, Output{word: word})
	}
	return outputStructSlice
}

func indicesToColor(mainStr, subStr string) map[int]bool {
	indices := make(map[int]bool)
	for i := 0; i <= len(mainStr)-len(subStr); i++ {
		if mainStr[i:i+len(subStr)] == subStr {
			indices = getIndicesToColor(i, i+len(subStr), indices)
			i += len(subStr) - 1
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
