package problem065

import "testing"

type inputType [][]int
type outputType []int

type testCase struct {
	input  inputType
	output outputType
}

func assert(t *testing.T, tCase testCase, f func([][]int) []int) bool {
	actual, expected := f(tCase.input), tCase.output
	for i, actualElement := range actual {
		if actualElement != expected[i] {
			t.Logf("actual: %v\nexpected: %v", actual, expected)
			t.Fail()
			return false
		}
	}
	return true
}

func TestTraverse(t *testing.T) {
	for _, tCase := range []testCase{
		{
			input: inputType{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
			output: outputType{1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12},
		},
	} {
		assert(t, tCase, TraverseMatrix)
	}
}
