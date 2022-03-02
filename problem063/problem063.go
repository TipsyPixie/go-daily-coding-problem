package problem063

type pair struct {
	first  bool
	second byte
}

func IncludesWord(letterMatrix [][]rune, word string) bool {
	wordLetters := []rune(word)
	collector := make(chan pair, len(letterMatrix))
	encodeAndCheck := func(letters []rune, wordLetter rune) {
		var encoded byte = 0
		included := true
		for i, letter := range letters {
			if letter != wordLetters[i] {
				included = false
			}
			if letter == wordLetter {
				encoded |= 1 << i
			}
		}
		collector <- pair{first: included, second: encoded}
	}
	for i, row := range letterMatrix {
		go encodeAndCheck(row, wordLetters[i])
	}

	mergedByte := ^byte(0)
	for i := 0; i < len(letterMatrix); i++ {
		select {
		case collectedPair := <-collector:
			if collectedPair.first {
				return true
			}
			mergedByte &= collectedPair.second
		}
	}
	return mergedByte > 0
}
