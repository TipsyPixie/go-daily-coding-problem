package problem064

import "testing"

type testCase struct {
	boardSize int
	result    int
}

func assert(t *testing.T, tCase testCase, f func(int) int) bool {
	actual, expected := f(tCase.boardSize), tCase.result
	if actual != expected {
		if isVerbose := testing.Verbose(); isVerbose {
			t.Logf("boardSize: %d", tCase.boardSize)
		}
		t.Logf("actual: %v, expected: %v", actual, expected)
		t.Fail()
		return false
	}
	return true
}

func TestCountKnightsPaths(t *testing.T) {
	for _, tCase := range []testCase{
		{
			boardSize: 1,
			result:    1,
		},
		{
			boardSize: 4,
			result:    0,
		},
		{
			boardSize: 5,
			result:    1728,
		},
	} {
		assert(t, tCase, CountKnightsPaths)
	}
}
