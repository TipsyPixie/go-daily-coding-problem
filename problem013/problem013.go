package problem013

func Run(input string, maxDistinctCount uint) uint {
	var longestSubstringSize uint = 0
	characterCount := make(map[string]int)
	substringStartsAt, substringEndsAt := 0, 0
	for substringEndsAt < len(input) {
		substringEndsWith := input[substringEndsAt : substringEndsAt+1]
		characterCount[substringEndsWith] += 1
		for len(characterCount) > int(maxDistinctCount) {
			substringStartsWith := input[substringStartsAt : substringStartsAt+1]
			characterCount[substringStartsWith] -= 1
			if characterCount[substringStartsWith] == 0 {
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
