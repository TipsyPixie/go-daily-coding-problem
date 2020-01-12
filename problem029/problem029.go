package problem029

import (
	"bytes"
	"strconv"
)

func Encode(plainText string) string {
	var lastCharacter rune
	var encodingBuffer bytes.Buffer
	runLength := 0
	for _, character := range []rune(plainText) {
		if character == lastCharacter {
			runLength++
		} else {
			if runLength > 0 {
				encodingBuffer.WriteString(strconv.Itoa(runLength))
				encodingBuffer.WriteRune(lastCharacter)
			}
			lastCharacter, runLength = character, 1
		}
	}

	if runLength > 0 {
		encodingBuffer.WriteString(strconv.Itoa(runLength))
		encodingBuffer.WriteRune(lastCharacter)
	}

	return encodingBuffer.String()
}
