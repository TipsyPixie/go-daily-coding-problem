package problem034

import (
	"fmt"
)

func minString(s1 string, s2 string) string {
	if s1 <= s2 {
		return s1
	}
	return s2
}

func ClosestPalindrome(word string) string {
	if len(word) <= 1 {
		return word
	}

	firstCharacter, lastCharacter := word[0], word[len(word)-1]
	if firstCharacter == lastCharacter {
		return fmt.Sprintf("%c%s%c", firstCharacter, ClosestPalindrome(word[1:len(word)-1]), lastCharacter)
	}

	candidate1 := fmt.Sprintf("%c%s%c", firstCharacter, ClosestPalindrome(word[1:]), firstCharacter)
	candidate2 := fmt.Sprintf("%c%s%c", lastCharacter, ClosestPalindrome(word[:len(word)-1]), lastCharacter)
	if len(candidate1) < len(candidate2) {
		return candidate1
	} else if len(candidate1) == len(candidate2) {
		return minString(candidate1, candidate2)
	} else {
		return candidate2
	}
}
