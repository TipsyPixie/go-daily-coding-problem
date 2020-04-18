package problem044

import (
	"fmt"
	"testing"
)

func TestCountInversions(t *testing.T) {
	testValues := [][]int{
		{2, 4, 1, 3, 5},
		{5, 4, 3, 2, 1},
	}
	correctAnswers := []int{3, 10}
	for i, testValue := range testValues {
		result := CountInversions(testValue)
		if result != correctAnswers[i] {
			t.Log(fmt.Sprintf("%v: %v != %v", testValue, result, correctAnswers[i]))
			t.Fail()
		}
	}
}
