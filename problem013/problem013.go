package problem013

func Run(input string, distinctCharactersLimit uint) uint {
	var longestSubstringSize uint = 0
	characterCount := make(map[rune]int)
	substringStartsAt, substringEndsAt := 0, 0
	for substringEndsAt < len(input) {
		startingCharacter := []rune(input)[substringEndsAt]
		characterCount[startingCharacter] += 1
		for uint(len(characterCount)) > distinctCharactersLimit {
			endingCharacter := []rune(input)[substringStartsAt]
			if characterCount[endingCharacter] > 1 {
				characterCount[endingCharacter]--
			} else {
				delete(characterCount, endingCharacter)
			}
			substringStartsAt++
		}
		if substringSize := uint(substringEndsAt - substringStartsAt + 1); substringSize > longestSubstringSize {
			longestSubstringSize = substringSize
		}
		substringEndsAt++
	}
	return longestSubstringSize
}
