package problem023

import "math"

type queue struct {
	innerSlice []boardState
}

type queueEmpty struct{}

func (err queueEmpty) Error() string {
	return "queue empty"
}

func newQueue(len int, cap int) *queue {
	return &queue{innerSlice: make([]boardState, len, cap)}
}

func (q *queue) enqueue(item boardState) {
	q.innerSlice = append(q.innerSlice, item)
}

func (q *queue) dequeue() (boardState, error) {
	if len(q.innerSlice) > 0 {
		item := q.innerSlice[0]
		q.innerSlice = q.innerSlice[1:]
		return item, nil
	} else {
		return boardState{}, queueEmpty{}
	}
}

type coordinate struct {
	column int
	row    int
}

type board [][]bool

func (thisBoard board) deepCopy() board {
	newBoard := make([][]bool, len(thisBoard), len(thisBoard))
	for i, column := range thisBoard {
		newBoard[i] = make([]bool, len(column), len(column))
		for j, row := range column {
			newBoard[i][j] = row
		}
	}
	return newBoard
}

type boardState struct {
	maze  board
	steps int
	start coordinate
	end   coordinate
}

func Run(maze [][]bool, startColumn int, startRow int, endColumn int, endRow int) int {
	q := newQueue(0, len(maze[0])*len(maze))
	q.enqueue(boardState{
		maze:  board(maze).deepCopy(),
		steps: 0,
		start: coordinate{column: startColumn, row: startRow},
		end:   coordinate{column: endColumn, row: endRow},
	})
	minimumSteps := math.MaxInt32
	for state, err := q.dequeue(); err == nil; state, err = q.dequeue() {
		switch {
		case state.steps >= len(state.maze)*len(state.maze[0]):
			fallthrough
		case state.start.column >= len(state.maze) || state.start.column < 0:
			fallthrough
		case state.start.row >= len(state.maze[state.start.column]) || state.start.row < 0:
			fallthrough
		case state.maze[state.start.column][state.start.row]:
			continue
		case state.start.column == state.end.column && state.start.row == state.end.row:
			if minimumSteps > state.steps {
				minimumSteps = state.steps
			}
		default:
			state.maze[state.start.column][state.start.row] = true
			state.steps++
			q.enqueue(boardState{
				maze:  state.maze.deepCopy(),
				steps: state.steps,
				start: coordinate{column: state.start.column - 1, row: state.start.row},
				end:   coordinate{column: state.end.column, row: state.end.row},
			})
			q.enqueue(boardState{
				maze:  state.maze.deepCopy(),
				steps: state.steps,
				start: coordinate{column: state.start.column, row: state.start.row - 1},
				end:   coordinate{column: state.end.column, row: state.end.row},
			})
			q.enqueue(boardState{
				maze:  state.maze.deepCopy(),
				steps: state.steps,
				start: coordinate{column: state.start.column, row: state.start.row + 1},
				end:   coordinate{column: state.end.column, row: state.end.row},
			})
			q.enqueue(boardState{
				maze:  state.maze.deepCopy(),
				steps: state.steps,
				start: coordinate{column: state.start.column + 1, row: state.start.row},
				end:   coordinate{column: state.end.column, row: state.end.row},
			})
		}
	}
	return minimumSteps
}
