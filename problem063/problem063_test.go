package problem063

import "testing"

type testCase struct {
	letterMatrix [][]rune
	word         string
	result       bool
}

func assert(t *testing.T, tCase testCase, f func([][]rune, string) bool) bool {
	actual, expected := f(tCase.letterMatrix, tCase.word), tCase.result
	if actual != expected {
		if isVerbose := testing.Verbose(); isVerbose {
			t.Log(tCase.letterMatrix)
		}
		t.Logf("actual: %v, expected: %v", actual, expected)
		t.Fail()
		return false
	}
	return true
}

func TestIncludesWord(t *testing.T) {
	for _, tCase := range []testCase{
		{
			letterMatrix: [][]rune{
				{'F', 'A', 'C', 'I'},
				{'O', 'B', 'Q', 'P'},
				{'A', 'N', 'O', 'B'},
				{'M', 'A', 'S', 'S'},
			},
			word:   "FOAM",
			result: true,
		},
		{
			letterMatrix: [][]rune{
				{'F', 'A', 'C', 'I'},
				{'O', 'B', 'Q', 'P'},
				{'A', 'N', 'O', 'B'},
				{'M', 'A', 'S', 'S'},
			},
			word:   "MASS",
			result: true,
		},
		{
			letterMatrix: [][]rune{
				{'F', 'A', 'C', 'I'},
				{'O', 'B', 'Q', 'P'},
				{'A', 'N', 'O', 'B'},
				{'M', 'A', 'S', 'S'},
			},
			word:   "ANOP",
			result: false,
		},
		{
			letterMatrix: [][]rune{
				{'F', 'A', 'C', 'I'},
				{'O', 'B', 'Q', 'P'},
				{'A', 'N', 'O', 'B'},
				{'M', 'A', 'S', 'S'},
			},
			word:   "IPBD",
			result: false,
		},
	} {
		assert(t, tCase, IncludesWord)
	}
}
