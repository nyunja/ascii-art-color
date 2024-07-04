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

			// Append all charaters(& their indices) to OutputStructSlice in event of '\n'
			if len(word) > 0 {
				outputStructSlice = append(outputStructSlice, Output{word: word})

				// Empty word for fresh population
				word = nil
			}

			// Handle multiple '\n' characters parsed with other characters
			if countNewlines > 1 {
				outputStructSlice = append(outputStructSlice, Output{newLine: "\n"})
			}

			// Print one new line for every '\n' parsed as input
			if countNewlines == len(s) {
				outputStructSlice = append(outputStructSlice, Output{newLine: "\n"})
				break
			}

			// Append individual characters and their indices to 'word' struct
		} else {
			word = append(word, Chars{index: i, character: ch})
			countNewlines = 0
		}
	}

	// Append any leftover characters that may be in word to OutputStructFile
	if len(word) > 0 {
		outputStructSlice = append(outputStructSlice, Output{word: word})
	}

	return outputStructSlice
}

// Returns map with the indices that need coloring
func indicesToColor(mainStr, subStr string) map[int]bool {
	indices := make(map[int]bool)

	// Generate map of indices whenever a section of mainStr matches subStr
	for i := 0; i <= len(mainStr)-len(subStr); i++ {
		if mainStr[i:i+len(subStr)] == subStr {
			indices = getIndicesToColor(i, i+len(subStr), indices)

			// Update i to find next match
			i += len(subStr) - 1
		}
	}

	return indices
}

// Populates map with indices that need coloring
func getIndicesToColor(i, j int, indices map[int]bool) map[int]bool {
	for i < j {
		indices[i] = true
		i++
	}

	return indices
}
