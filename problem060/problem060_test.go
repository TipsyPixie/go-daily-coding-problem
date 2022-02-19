package problem060

import "testing"

type argumentSet struct {
	numbers []int
	result  bool
}

func assert(t *testing.T, actual bool, expected bool) bool {
	if actual != expected {
		t.Logf("Actual: %t / Expected: %t", actual, expected)
		return false
	}
	return true
}

func TestCanBePartitioned(t *testing.T) {
	testCases := []argumentSet{
		{
			numbers: []int{15, 5, 20, 10, 35, 15, 10},
			result:  true,
		},
		{
			numbers: []int{15, 5, 20, 10, 35},
			result:  false,
		},
		{
			numbers: []int{15, 7, 8},
			result:  true,
		},
		{
			numbers: []int{},
			result:  false,
		},
		{
			numbers: []int{11, 10, 1},
			result:  true,
		},
		{
			numbers: []int{4},
			result:  false,
		},
	}

	for _, testCase := range testCases {
		if !assert(t, CanBePartitioned(testCase.numbers), testCase.result) {
			t.FailNow()
		}
	}
}
