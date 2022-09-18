package problem013

func Run(input string, distinctCharactersLimit uint) uint {
	var longestSubstringSize uint = 0
	characterCount := make(map[rune]int)
	substringStartsAt, substringEndsAt := 0, 0
	for substringEndsAt < len(input) {
		startingCharacter := []rune(input)[substringEndsAt]
		characterCount[startingCharacter] += 1
		for uint(len(characterCount)) > distinctCharactersLimit {
			startingCharacter := []rune(input)[substringStartsAt]
			if characterCount[startingCharacter] > 1 {
				characterCount[startingCharacter]--
			} else {
				delete(characterCount, startingCharacter)
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
