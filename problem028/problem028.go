package problem028

import "bytes"

func Justify(words []string, lineLength int) []string {
	justifiedWords, lineSizeLeft, wordsInline, spaceCounts := initValues(words, lineLength)
	for wordIndex := 0; wordIndex < len(words); {
		word := words[wordIndex]
		if (len(wordsInline) == 0 && len(word) <= lineSizeLeft) || (len(wordsInline) > 0 && len(word)+1 <= lineSizeLeft) {
			if len(wordsInline) > 0 {
				spaceCounts = append(spaceCounts, 1)
				lineSizeLeft--
			}
			wordsInline = append(wordsInline, word)
			lineSizeLeft -= len(word)
			wordIndex++
		} else {
			justifiedWords = append(justifiedWords, justifyLine(lineSizeLeft, wordsInline, spaceCounts))
			_, lineSizeLeft, wordsInline, spaceCounts = initValues(words, lineLength)
		}
	}
	justifiedWords = append(justifiedWords, justifyLine(lineSizeLeft, wordsInline, spaceCounts))

	return justifiedWords
}

func justifyLine(lineSizeLeft int, wordsInline []string, spaceCounts []int) string {
	if len(wordsInline) == 1 {
		spaceCounts = append(spaceCounts, 0)
	}
	for lineSize := 0; lineSize < lineSizeLeft; lineSize++ {
		spaceCounts[lineSize%len(spaceCounts)]++
	}
	if len(wordsInline) > 1 {
		spaceCounts = append(spaceCounts, 0)
	}

	var buffer bytes.Buffer
	for inlineIndex, word := range wordsInline {
		buffer.WriteString(word)
		for spaceCount := 0; spaceCount < spaceCounts[inlineIndex]; spaceCount++ {
			buffer.WriteString(" ")
		}
	}

	return buffer.String()
}

func initValues(words []string, lineLength int) ([]string, int, []string, []int) {
	return make([]string, 0, len(words)), lineLength, make([]string, 0, len(words)), make([]int, 0, len(words)/2)
}
