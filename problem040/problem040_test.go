package problem040

import (
	"fmt"
	"testing"
)

func TestFindSingleNumber(t *testing.T) {
	values := [][]int{
		{6, 6, 3, 1, 3, 6, 3},
		{13, 19, 13, 13},
		{-14, -14, 22, -14, 22, -2, -2, -15, -15, -15, -2, 22, 0},
	}
	correctAnswers := []int{1, 19, 0}
	for i := range values {
		if answer := FindSingleNumber(values[i]); answer != correctAnswers[i] {
			t.Log(fmt.Sprintf("%d != %d", answer, correctAnswers[i]))
			t.FailNow()
		}
	}
}
