package problem035

func partialSegregate(letter rune, in []rune) int {
	leftIndex, rightIndex := 0, len(in)
	for leftIndex < rightIndex {
		leftLetter := in[leftIndex]
		if leftLetter == letter {
			leftIndex++
		} else if rightIndex == len(in) {
			rightIndex--
		} else if rightLetter := in[rightIndex]; rightLetter != letter {
			rightIndex--
		} else {
			in[leftIndex], in[rightIndex] = rightLetter, leftLetter
		}
	}
	return leftIndex + 1
}

func Segregate(letters []rune) {
	letterTypes := []rune{'R', 'G', 'B'}
	for typeIndex, startingIndex := 0, 0; typeIndex < len(letterTypes) && startingIndex < len(letters); typeIndex++ {
		startingIndex = partialSegregate(letterTypes[typeIndex], letters[startingIndex:])
	}
}
