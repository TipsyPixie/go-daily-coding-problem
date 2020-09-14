package problem054

// assume a valid solution exists

type game struct {
	numberAt  [][]int
	rowSet    []map[int]bool
	columnSet []map[int]bool
	boxSet    [][]map[int]bool
}

const empty = 0

func (thisGame *game) set(row int, column int, number int) *game {
	if number == empty {
		delete(thisGame.rowSet[row], thisGame.numberAt[row][column])
		delete(thisGame.columnSet[column], thisGame.numberAt[row][column])
		delete(thisGame.boxSet[row/3][column/3], thisGame.numberAt[row][column])
	} else {
		thisGame.rowSet[row][number] = true
		thisGame.columnSet[column][number] = true
		thisGame.boxSet[row/3][column/3][number] = true
	}
	thisGame.numberAt[row][column] = number
	return thisGame
}

func (thisGame *game) resolveAt(row int, column int) bool {
	if row > 8 {
		return true
	}

	nextRow, nextColumn := row, column+1
	if nextColumn > 8 {
		nextRow, nextColumn = row+1, 0
	}

	if thisGame.numberAt[row][column] != empty {
		return thisGame.resolveAt(nextRow, nextColumn)
	}

	for i := 1; i < 10; i++ {
		alreadyInBox := thisGame.boxSet[row/3][column/3][i]
		alreadyOnRow := thisGame.rowSet[row][i]
		alreadyOnColumn := thisGame.columnSet[column][i]
		if !alreadyInBox && !alreadyOnRow && !alreadyOnColumn {
			thisGame.set(row, column, i)
			if thisGame.resolveAt(nextRow, nextColumn) {
				return true
			}
			thisGame.set(row, column, empty)
		}
	}
	return false
}

func ResolveSudoku(values [][]int) [][]int {
	newGame := &game{
		numberAt:  make([][]int, 9, 9),
		rowSet:    make([]map[int]bool, 9, 9),
		columnSet: make([]map[int]bool, 9, 9),
		boxSet:    make([][]map[int]bool, 3, 3),
	}
	for i := range newGame.numberAt {
		newGame.numberAt[i] = make([]int, 9, 9)
	}
	for i := range newGame.rowSet {
		newGame.rowSet[i] = map[int]bool{}
		newGame.columnSet[i] = map[int]bool{}
	}
	for i := range newGame.boxSet {
		newGame.boxSet[i] = make([]map[int]bool, 3, 3)
		for j := range newGame.boxSet[i] {
			newGame.boxSet[i][j] = map[int]bool{}
		}
	}

	for rowIndex, row := range values {
		for columnIndex := range row {
			newGame.set(rowIndex, columnIndex, values[rowIndex][columnIndex])
		}
	}
	newGame.resolveAt(0, 0)
	return newGame.numberAt
}
