package problem058

import (
	"testing"
)

type argumentSet struct {
	numbers      []int
	targetNumber int
	result       *int
}

func assertValue(t *testing.T, actual *int, expected *int) bool {
	if (actual != nil || expected != nil) && (*actual != *expected) {
		t.Logf("Actual: %d / Expected: %d", *actual, *expected)
		return false
	}
	return true
}

func TestSearch(t *testing.T) {
	testCases := []argumentSet{
		{
			numbers:      []int{13, 18, 25, 2, 8, 10},
			targetNumber: 8,
			result:       getIntPointer(4),
		},
		{
			numbers:      []int{13, 18, 25, 2, 8, 10},
			targetNumber: 7,
			result:       nil,
		},
		{
			numbers:      []int{2, 8, 10, 11, 13, 18, 25},
			targetNumber: 1,
			result:       nil,
		},
		{
			numbers:      []int{},
			targetNumber: 8,
			result:       nil,
		},
	}

	for _, testCase := range testCases {
		if !assertValue(t, Search(testCase.numbers, testCase.targetNumber), testCase.result) {
			t.FailNow()
		}
	}
}
