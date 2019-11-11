package problem017

import "testing"

func TestRun(t *testing.T) {
	arg := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	if result := Run(arg); result != 32 {
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	s := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	result := Run(s)
	if result != 20 {
		t.FailNow()
	}
}
