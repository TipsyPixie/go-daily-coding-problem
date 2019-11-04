package problem2

import (
	"reflect"
	"testing"
)

func TestRunSimple(t *testing.T) {
	if !reflect.DeepEqual(Run([]int{1, 2, 3, 4, 5}), []int{120, 60, 40, 30, 24}) {
		t.Fail()
	}
	if !reflect.DeepEqual(Run([]int{3, 2, 1}), []int{2, 3, 6}) {
		t.Fail()
	}
}

func TestRunLong(t *testing.T) {
	size := 50
	production := 1
	args := make([]int, size, size)
	for i := 0; i < size; i++ {
		args[i] = i%3 + 1
		production *= i%3 + 1
	}
	expectedResult := make([]int, size, size)
	for i, value := range args {
		expectedResult[i] = production / value
	}
	if !reflect.DeepEqual(Run(args), expectedResult) {
		t.Fail()
	}
}
