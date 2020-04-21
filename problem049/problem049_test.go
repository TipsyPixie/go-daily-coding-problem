package problem049

import (
	"fmt"
	"testing"
)

func TestMaximizeContiguousSum(t *testing.T) {
	testArgs := [][]int{
		{34, -50, 42, 14, -5, 86},
		{-5, -1, -8, -9},
	}
	correctResult := []int{137, 0}
	for i, arg := range testArgs {
		result := MaximizeContiguousSum(arg)
		if result != correctResult[i] {
			t.Log(fmt.Sprintf("%v: %d != %d", arg, result, correctResult))
			t.Fail()
		}
	}
}
