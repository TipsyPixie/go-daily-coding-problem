package problem021

import "testing"

func TestRun(t *testing.T) {
	lectures := []Schedule{NewLecture(30, 75), NewLecture(0, 50), NewLecture(60, 150)}
	if Run(lectures) != 2 {
		t.FailNow()
	}
}

//[(30, 75), (0, 50), (60, 150)],
