package problem013

func Run(input string, distinctCharactersLimit uint) uint {
	var longestSubstringSize uint = 0
	characterCount := make(map[rune]int)
	substringStartsAt, substringEndsAt := 0, 0
	for substringEndsAt < len(input) {
		substringEndsWith := []rune(input)[substringEndsAt]
		characterCount[substringEndsWith] += 1
		for uint(len(characterCount)) > distinctCharactersLimit {
			substringStartsWith := []rune(input)[substringStartsAt]
			if characterCount[substringStartsWith] > 1 {
				characterCount[substringStartsWith] -= 1
			} else {
				delete(characterCount, substringStartsWith)
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
