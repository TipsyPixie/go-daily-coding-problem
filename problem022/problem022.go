package problem022

// use either of strings.LastIndex() or your own pattern matching function

func getFailoverTable(pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	patternCharacters := []rune(pattern)
	prefixLengths := make([]int, len(patternCharacters), len(patternCharacters))
	prefixLengths[0] = 0
	for i := 1; i < len(patternCharacters); i++ {
		lastPrefixLength := prefixLengths[i-1]
		for lastPrefixLength > 0 && patternCharacters[lastPrefixLength] != patternCharacters[i] {
			lastPrefixLength = prefixLengths[lastPrefixLength-1]
		}
		if patternCharacters[i] == patternCharacters[lastPrefixLength] {
			prefixLengths[i] = lastPrefixLength + 1
		} else {
			prefixLengths[i] = 0
		}
	}
	return prefixLengths
}

func KMPMatch(pattern string, from string) int {
	failoverTable := getFailoverTable(pattern)
	for patternStartsAt, patternIndex := 0, 0; patternStartsAt+patternIndex < len(from); {
		fromIndex := patternStartsAt + patternIndex
		if []rune(pattern)[patternIndex] == []rune(from)[fromIndex] {
			patternIndex++
			if patternIndex >= len(pattern) {
				return patternStartsAt
			}
		} else {
			if patternIndex > 0 && failoverTable[patternIndex-1] > 0 {
				patternStartsAt += patternIndex - failoverTable[patternIndex-1]
				patternIndex -= patternIndex - failoverTable[patternIndex-1]
			} else {
				patternStartsAt += patternIndex + 1
				patternIndex = 0
			}
		}
	}
	return -1
}

func Run(words []string, sentence string) []string {
	wordsFound := make([]string, 0, len(words))
	for _, word := range words {
		index := KMPMatch(word, sentence)
		if index >= 0 {
			wordsFound = append(wordsFound, Run(words, sentence[:index])...)
			wordsFound = append(wordsFound, word)
			wordsFound = append(wordsFound, Run(words, sentence[index+len(word):])...)
			return wordsFound
		}
	}
	return wordsFound
}
