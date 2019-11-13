package problem019

import "testing"

func TestRun(t *testing.T) {
	result := Run([][]int{
		{5, 8, 6},
		{19, 14, 13},
		{7, 5, 12},
		{14, 15, 17},
		{3, 20, 10},
	})
	if result != 43 {
		t.FailNow()
	}
}
