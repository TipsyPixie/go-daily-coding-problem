package problem046

import (
	"fmt"
	"testing"
)

func TestGetPalindromicSubstring(t *testing.T) {
	testValues := []string{
		"aabcdcb",
		"bananas",
		"foxjump",
	}
	correctAnswers := []string{
		"bcdcb",
		"anana",
		"p",
	}
	for i, testValue := range testValues {
		answer := GetPalindromicSubstring(testValue)
		if answer != correctAnswers[i] {
			t.Log(fmt.Sprintf("%s: %s != %s", testValue, answer, correctAnswers[i]))
			t.Fail()
		}
	}
}
