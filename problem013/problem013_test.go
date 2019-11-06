package problem013

import "testing"

func TestRun(t *testing.T) {
	if Run("abcba", 2) != 3 {
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	if Run("abbcbcbbac", 2) != 7 {
		t.FailNow()
	}
}
