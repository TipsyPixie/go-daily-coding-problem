package problem021

import "testing"

func TestRun(t *testing.T) {
	lectures := []Schedule{NewLecture(30, 75), NewLecture(0, 50), NewLecture(60, 150)}
	result := Run(lectures)
	if result != 2 {
		t.Log(result)
		t.FailNow()
	}
}
