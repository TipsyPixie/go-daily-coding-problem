package problem012

import "testing"

func TestRun(t *testing.T) {
	if Run(4, []int{1, 2}) != 5 {
		t.FailNow()
	}
}
