package problem001

import "testing"

func TestRunSimple(t *testing.T) {
	result := Run([]int{10, 15, 3, 7}, 17)
	if !result {
		t.FailNow()
	}
}

func TestRunLong(t *testing.T) {
	size := 1000000
	args := make([]int, size, size)
	for i := 0; i < size; i++ {
		args[i] = i
	}
	result := Run(args, size*2)
	if result {
		t.FailNow()
	}
}
