package problem043

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(32).Push(4).Push(10).Push(12).Push(4).Push(12)

	correctAnswer := 12
	if maxValue, err := stack.Max(); err != nil || maxValue != correctAnswer {
		t.Log(fmt.Sprintf("Max %d != %d", maxValue, correctAnswer), err)
		t.Fail()
	}

	correctAnswer = 12
	if poppedValue, err := stack.Pop(); err != nil || poppedValue != correctAnswer {
		t.Log(fmt.Sprintf("Pop %d != %d", poppedValue, correctAnswer), err)
		t.Fail()
	}

	correctAnswer = 12
	if maxValue, err := stack.Max(); err != nil || maxValue != correctAnswer {
		t.Log(fmt.Sprintf("Max %d != %d", maxValue, correctAnswer), err)
		t.Fail()
	}

	correctAnswer = 4
	if poppedValue, err := stack.Pop(); err != nil || poppedValue != correctAnswer {
		t.Log(fmt.Sprintf("Pop %d != %d", poppedValue, correctAnswer), err)
		t.Fail()
	}

	_, _ = stack.Pop()

	correctAnswer = 10
	if maxValue, err := stack.Max(); err != nil || maxValue != correctAnswer {
		t.Log(fmt.Sprintf("Max %d != %d", maxValue, correctAnswer), err)
		t.Fail()
	}

	_, _ = stack.Pop()

	correctAnswer = 4
	if maxValue, err := stack.Max(); err != nil || maxValue != correctAnswer {
		t.Log(fmt.Sprintf("Max %d != %d", maxValue, correctAnswer), err)
		t.Fail()
	}

	_, _ = stack.Pop()

	if _, err := stack.Pop(); err == nil {
		t.Log("Pop without error")
		t.Fail()
	}
}
