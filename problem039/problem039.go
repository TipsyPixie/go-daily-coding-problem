package problem039

import (
	"errors"
	"fmt"
)

type culuturePlate [][]int

const (
	maxSize = 512
	dead    = 0
	live    = 1
)

func NewCuluturePlate(maxRow int, maxColumn int) *culuturePlate {
	plate := make(culuturePlate, maxRow+1, maxRow+1)
	for i := range plate {
		plate[i] = make([]int, maxColumn+1, maxColumn+1)
	}
	return &plate
}

func (thisPlate *culuturePlate) copy() *culuturePlate {
	plate := make(culuturePlate, len(*thisPlate), len(*thisPlate))
	for rowIndex := range plate {
		plate[rowIndex] = make([]int, len((*thisPlate)[rowIndex]), len((*thisPlate)[rowIndex]))
	}
	return &plate
}

func (thisPlate *culuturePlate) get(row int, column int) int {
	if row < 0 || row >= len(*thisPlate) || column < 0 || column >= len((*thisPlate)[row]) {
		return dead
	}
	return (*thisPlate)[row][column]
}

func (thisPlate *culuturePlate) set(row int, column int, state int) {
	(*thisPlate)[row][column] = state
}

func (thisPlate *culuturePlate) countNeighbors(row int, column int) int {
	neighborCoords := [][]int{
		{row - 1, column - 1},
		{row - 1, column},
		{row - 1, column + 1},
		{row, column - 1},
		{row, column + 1},
		{row + 1, column - 1},
		{row + 1, column},
		{row + 1, column + 1},
	}
	neighborCount := 0
	for _, coordinates := range neighborCoords {
		if thisPlate.get(coordinates[0], coordinates[1]) == live {
			neighborCount++
		}
	}
	return neighborCount
}

func (thisPlate *culuturePlate) getNextCellState(currentState int, neighborCount int) int {
	if currentState == live {
		switch {
		case neighborCount < 2:
			return dead
		case neighborCount == 2 || neighborCount == 3:
			return live
		case neighborCount > 3:
			return dead
		}
	} else {
		if neighborCount == 3 {
			return live
		}
	}
	return currentState
}

func (thisPlate *culuturePlate) next() (*culuturePlate, error) {
	nextPlate := thisPlate.copy()
	for rowIndex := 0; rowIndex < len(*thisPlate)+1; rowIndex++ {
		rowLength := 0
		if rowIndex < len(*thisPlate) {
			rowLength = len((*thisPlate)[rowIndex])
		} else {
			rowLength = len((*thisPlate)[rowIndex - 1])
		}
		for columnIndex := 0; columnIndex < rowLength+1; columnIndex++ {
			neighborCount := thisPlate.countNeighbors(rowIndex, columnIndex)
			nextCellState := thisPlate.getNextCellState(thisPlate.get(rowIndex, columnIndex), neighborCount)

			if rowIndex < len(*thisPlate) && columnIndex < len((*thisPlate)[rowIndex]) {
				nextPlate.set(rowIndex, columnIndex, nextCellState)
			} else if nextCellState == live {
				if rowIndex >= maxSize || columnIndex >= maxSize {
					return nil, errors.New("max size exceeded")
				}
				if rowIndex >= len(*nextPlate) {
					var newRow []int
					*nextPlate = append(*nextPlate, newRow)
				}
				if columnIndex >= len((*nextPlate)[rowIndex]) {
					difference := columnIndex - len((*nextPlate)[rowIndex]) + 1
					(*nextPlate)[rowIndex] = append((*nextPlate)[rowIndex], make([]int, difference, difference)...)
				}
				nextPlate.set(rowIndex, columnIndex, nextCellState)
			}
		}
	}
	return nextPlate, nil
}

func GameOfLife(liveCellsCoords [][]int, steps int) error {
	maxRow, maxColumn := 0, 0
	for _, coords := range liveCellsCoords {
		if coords[0] > maxRow {
			maxRow = coords[0]
		}
		if coords[1] > maxColumn {
			maxColumn = coords[1]
		}
	}

	plate := NewCuluturePlate(maxRow, maxColumn)
	for _, coords := range liveCellsCoords {
		plate.set(coords[0], coords[1], live)
	}

	for i := 0; i < steps+1; i++ {
		if i < steps {
			next, err := plate.next()
			if err != nil {
				return err
			}
			plate = next
		}
	}

	for _, row := range *plate {
		for _, cell := range row {
			if cell == live {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}
