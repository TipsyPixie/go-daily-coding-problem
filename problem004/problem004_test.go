package problem004

import (
	"math/rand"
	"testing"
	"time"
)

func TestRunSimple(t *testing.T) {
	args, expectedResult := []int{3, 4, -1, 1}, 2
	if Run(args) != expectedResult {
		t.Fail()
	}

	args, expectedResult = []int{1, 2, 0}, 3
	if Run(args) != expectedResult {
		t.Fail()
	}
}

func TestRunLong(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	size := 1000000
	existingNumbers := make(map[int]bool)
	args := make([]int, size, size)
	for index := range args {
		args[index] = rand.Intn(size+1) - size/2
		if args[index] > 0 {
			existingNumbers[args[index]] = true
		}
	}
	expectedResult := 1
	for expectedResult <= len(args) && existingNumbers[expectedResult] {
		expectedResult++
	}
	if Run(args) != expectedResult {
		t.Fail()
	}
}
