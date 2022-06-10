package problem065

type coord struct {
	row int
	col int
}

func traverseMatrix(matrix [][]int) []int {
	visited := map[coord]bool{}
	var result []int
	rowDiffs := []int{1, 0, -1, 0}
	colDiffs := []int{0, 1, 0, -1}

	diffIndex := 0
	row, col := -1, 0
	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		rowNext, colNext := row+rowDiffs[diffIndex], col+colDiffs[diffIndex]
		if rowNext >= 0 && rowNext < len(matrix[0]) && colNext >= 0 && colNext < len(matrix) && !visited[coord{row: rowNext, col: colNext}] {
			row, col = rowNext, colNext
			result = append(result, matrix[col][row])
			visited[coord{row: row, col: col}] = true
		} else {
			diffIndex = (diffIndex + 1) % 4
		}
	}

	return result
}
