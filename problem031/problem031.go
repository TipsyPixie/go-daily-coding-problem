package problem031

func minInt(nums ...int) int {
	minimum := nums[0]
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func MeasureEditDistance(value1 string, value2 string) int {
	// make sure width > height of the matrix
	if len(value1) < len(value2) {
		value1, value2 = value2, value1
	}

	matrixWidth, matrixHeight := len(value1)+1, len(value2)+1
	distanceMatrix := make([][]int, matrixHeight, matrixHeight)
	for rowIndex := range distanceMatrix {
		distanceMatrix[rowIndex] = make([]int, matrixWidth, matrixWidth)
	}

	for initialRow, initialColumn := 0, 0; initialRow < matrixHeight && initialColumn < matrixWidth; {
		for row, column := initialRow, initialColumn; 0 <= row && row < matrixHeight && 0 <= column && column < matrixWidth; {
			switch {
			case row == 0:
				distanceMatrix[row][column] = column
			case column == 0:
				distanceMatrix[row][column] = row
			default:
				if value2[row-1] == value1[column-1] {
					distanceMatrix[row][column] = distanceMatrix[row-1][column-1]
				} else {
					distanceMatrix[row][column] = 1 + minInt(
						distanceMatrix[row-1][column],
						distanceMatrix[row-1][column-1],
						distanceMatrix[row][column-1],
					)
				}
			}
			row++
			column--
		}

		if initialColumn < matrixWidth-1 {
			initialColumn++
		} else {
			initialRow++
		}
	}

	return distanceMatrix[matrixHeight-1][matrixWidth-1]
}
