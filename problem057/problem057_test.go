package problem057

import (
	"reflect"
	"testing"
)

type argumentSet struct {
	fullSentence string
	lengthLimit  int
	result       []string
}

func TestSplitLine(t *testing.T) {
	testCases := []argumentSet{
		{
			fullSentence: "the quick brown fox jumps over the lazy dog",
			lengthLimit:  10,
			result:       []string{"the quick", "brown fox", "jumps over", "the lazy", "dog"},
		},
		{
			fullSentence: "the quick brown fox jumps over the lazy dog",
			lengthLimit:  1,
			result:       nil,
		},
	}

	for _, testCase := range testCases {
		actual, expected := SplitLine(testCase.fullSentence, testCase.lengthLimit), testCase.result
		if actual == nil && expected == nil {
			t.Logf("Actual %s != Expected %s", "nil", "nil")
		} else {
			t.Logf("Actual %s != Expected %s", actual, expected)
		}
		if (actual == nil && expected == nil) && !reflect.DeepEqual(actual, expected) {
			t.Fail()
		}
	}
}
