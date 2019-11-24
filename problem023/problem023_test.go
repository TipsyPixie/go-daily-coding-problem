package problem023

import (
	"testing"
)

func TestRun(t *testing.T) {
	maze := [][]bool{
		{false, false, false, false},
		{true, true, false, true},
		{false, false, false, false},
		{false, false, false, false},
	}
	result := Run(maze, 3, 0, 0, 0)
	if result != 7 {
		t.Log(result)
		t.FailNow()
	}
}
