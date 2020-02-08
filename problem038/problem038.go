package problem038

type chessBoard [][]bool

func newChessBoard(size int) *chessBoard {
	board := make(chessBoard, size, size)
	for i := range board {
		board[i] = make([]bool, size, size)
	}
	return &board
}

func (thisBoard *chessBoard) hasQueen(row int, column int) bool {
	return (*thisBoard)[row][column]
}

func (thisBoard *chessBoard) putQueen(row int, column int) {
	(*thisBoard)[row][column] = true
}

func (thisBoard *chessBoard) removeQueen(row int, column int) {
	(*thisBoard)[row][column] = false
}

func (thisBoard *chessBoard) isSafe(row int, column int) bool {
	onBoard := func(r int, c int) bool {
		return r >= 0 && r < len(*thisBoard) && c >= 0 && c < len(*thisBoard)
	}

	// horizontal check is unnecessary

	// vertical
	for otherRow := 0; onBoard(otherRow, column); otherRow++ {
		if thisBoard.hasQueen(otherRow, column) {
			return false
		}
	}

	// diagonal
	for otherRow, otherColumn := row, column; onBoard(otherRow, otherColumn); otherRow, otherColumn = otherRow+1, otherColumn+1 {
		if thisBoard.hasQueen(otherRow, otherColumn) {
			return false
		}
	}
	for otherRow, otherColumn := row, column; onBoard(otherRow, otherColumn); otherRow, otherColumn = otherRow-1, otherColumn-1 {
		if thisBoard.hasQueen(otherRow, otherColumn) {
			return false
		}
	}
	for otherRow, otherColumn := row, column; onBoard(otherRow, otherColumn); otherRow, otherColumn = otherRow+1, otherColumn-1 {
		if thisBoard.hasQueen(otherRow, otherColumn) {
			return false
		}
	}
	for otherRow, otherColumn := row, column; onBoard(otherRow, otherColumn); otherRow, otherColumn = otherRow-1, otherColumn+1 {
		if thisBoard.hasQueen(otherRow, otherColumn) {
			return false
		}
	}

	return true
}

func arrange(boardState *chessBoard, alreadyDeployed int) int {
	boardSize := len(*boardState)
	if alreadyDeployed == boardSize {
		return 1
	}

	arrangementCount := 0
	for column := 0; column < boardSize; column++ {
		if boardState.isSafe(alreadyDeployed, column) {
			boardState.putQueen(alreadyDeployed, column)
			arrangementCount += arrange(boardState, alreadyDeployed+1)
			boardState.removeQueen(alreadyDeployed, column)
		}
	}
	return arrangementCount
}

func GetArrangementsCount(boardSize int) int {
	board := newChessBoard(boardSize)
	return arrange(board, 0)
}
