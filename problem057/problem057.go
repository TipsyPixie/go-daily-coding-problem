package problem057

import (
	"strings"
)

func SplitLine(fullSentence string, lengthLimit int) []string {
	lines := make([]string, 0)
	currentLineWords, currentLineLength := make([]string, 0, lengthLimit), 0
	for _, word := range strings.Split(fullSentence, " ") {
		if len(word) > lengthLimit {
			return nil
		} else if lineLengthToBe := currentLineLength + 1 + len(word); lineLengthToBe <= lengthLimit {
			currentLineWords, currentLineLength = append(currentLineWords, word), currentLineLength+1+len(word)
		} else {
			lines = append(lines, strings.Join(currentLineWords, " "))
			currentLineWords, currentLineLength = []string{word}, len(word)
		}
	}

	if currentLineLength > 0 {
		lines = append(lines, strings.Join(currentLineWords, " "))
	}

	return lines
}
