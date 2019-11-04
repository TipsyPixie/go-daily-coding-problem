package problem7

import "testing"

func TestRun(t *testing.T) {
	if Run("111") != 3 {
		t.FailNow()
	}
	if Run("1111") != 5 {
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	if Run("10") != 1 {
		t.FailNow()
	}
	if Run("101") != 1 {
		t.FailNow()
	}
	if Run("1010") != 1 {
		t.FailNow()
	}
	if Run("101010") != 1 {
		t.FailNow()
	}
}
