package problem009

import "testing"

func TestRun(t *testing.T) {
	args := []int{2, 4, 6, 2, 5}
	expected := 13
	if Run(args) != expected {
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	args := []int{5, 1, 1, 5}
	expected := 10
	if Run(args) != expected {
		t.FailNow()
	}
}

func TestRun3(t *testing.T) {
	args := []int{10, -1, -1, 0, -7, -1, 0, 3, 7}
	expected := 17
	if Run(args) != expected {
		t.FailNow()
	}
}
