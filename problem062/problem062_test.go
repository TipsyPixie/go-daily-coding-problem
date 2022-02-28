package problem062

import "testing"

type argumentSet struct {
	rowCount    int
	columnCount int
	result      int
}

func assert(t *testing.T, actual int, expected int) bool {
	if actual != expected {
		t.Logf("Actual: %v / Expected: %v", actual, expected)
		return false
	}
	return true
}

func TestCountWay(t *testing.T) {
	testCases := []argumentSet{
		{
			rowCount:    2,
			columnCount: 2,
			result:      2,
		},
		{
			rowCount:    5,
			columnCount: 5,
			result:      70,
		},
	}

	for _, testCase := range testCases {
		if !assert(t, CountWay(testCase.rowCount, testCase.columnCount), testCase.result) {
			t.FailNow()
		}
	}
}
