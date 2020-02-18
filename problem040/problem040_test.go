package problem039

import "testing"

func TestGameOfLife(t *testing.T) {
	err := GameOfLife([][]int{
		{0, 0}, {0, 2}, {0, 3}, {1, 5},
		{1, 1}, {1, 4},
		{2, 2}, {2, 3}, {2, 4},
	}, 50)

	if err != nil {
		t.FailNow()
	}
}
