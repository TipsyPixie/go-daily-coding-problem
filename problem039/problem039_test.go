package problem039

import (
	"fmt"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	result, err := GameOfLife([][]int{
		{0, 0}, {0, 2}, {0, 3}, {1, 5},
		{1, 1}, {1, 4},
		{2, 2}, {2, 3}, {2, 4},
	}, 50)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	correctAnswer := ".......\n.......##\n......#..#\n.......##\n.......\n.......\n...\n"
	if result != correctAnswer {
		t.Log(fmt.Sprintf("--- expected ---\n%s\n--- actual ---\n%s\n", correctAnswer, result))
	}
}
