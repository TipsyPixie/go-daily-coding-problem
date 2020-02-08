package problem038

import (
	"fmt"
	"testing"
)

func TestGetArrangementsCount(t *testing.T) {
	correctAnswers := []int{1, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680, 14200, 73712}
	for i := range correctAnswers {
		answer := GetArrangementsCount(i)
		if answer != correctAnswers[i] {
			t.Log(fmt.Sprintf("%d: %d != %d", i, answer, correctAnswers[i]))
			t.FailNow()
		}
	}
}
