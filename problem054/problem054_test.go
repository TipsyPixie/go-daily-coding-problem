package problem054

import (
	"testing"
)

func TestResolveSudoku(t *testing.T) {
	sudoku := [][]int{
		{8, 0, 0, 0, 1, 0, 0, 7, 0},
		{0, 0, 4, 0, 5, 0, 0, 0, 2},
		{0, 0, 9, 0, 0, 4, 0, 3, 0},
		{0, 5, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 6, 8, 0},
		{9, 0, 0, 0, 6, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 3, 4, 0},
		{7, 0, 0, 5, 2, 0, 0, 0, 1},
		{0, 8, 0, 0, 0, 0, 0, 0, 0},
	}
	correctAnswer := [][]int{
		{8, 6, 5, 3, 1, 2, 4, 7, 9},
		{3, 7, 4, 9, 5, 8, 1, 6, 2},
		{1, 2, 9, 6, 7, 4, 5, 3, 8},
		{6, 5, 8, 2, 4, 7, 9, 1, 3},
		{4, 3, 2, 1, 9, 5, 6, 8, 7},
		{9, 1, 7, 8, 6, 3, 2, 5, 4},
		{2, 9, 6, 7, 8, 1, 3, 4, 5},
		{7, 4, 3, 5, 2, 6, 8, 9, 1},
		{5, 8, 1, 4, 3, 9, 7, 2, 6},
	}
	for rowIndex, row := range ResolveSudoku(sudoku) {
		for columnIndex, value := range row {
			if value != correctAnswer[rowIndex][columnIndex] {
				t.Logf("%d != %d at (%d, %d)", value, correctAnswer[rowIndex][columnIndex], rowIndex, columnIndex)
				t.Fail()
			}
		}
	}
}
