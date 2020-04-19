package problem047

import (
	"fmt"
	"testing"
)

func TestMaximizeProfit(t *testing.T) {
	testValue := []int{9, 11, 8, 5, 7, 10}
	answer := MaximizeProfit(testValue)
	correctAnswer := 5
	if answer != correctAnswer {
		t.Log(fmt.Sprintf("%v: %d != %d", testValue, answer, correctAnswer))
		t.Fail()
	}
}
