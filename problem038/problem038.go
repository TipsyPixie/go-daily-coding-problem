package problem038

func arrange(rowSize int, alreadyDeployed []int) int {
	if len(alreadyDeployed) == rowSize {
		return 1
	}

	onDiagonal := func(r1 int, c1 int, r2 int, c2 int) bool {
		return c1-r1 == c2-r2 || c1+r1 == c2+r2
	}

	count := 0
	row := len(alreadyDeployed)
	for column := 0; column < rowSize; column++ {
		safeToDeploy := true
		for filledRow, filledColumn := range alreadyDeployed {
			if filledColumn == column || onDiagonal(row, column, filledRow, filledColumn) {
				safeToDeploy = false
				break
			}
		}
		if safeToDeploy {
			count += arrange(rowSize, append(alreadyDeployed, column))
		}
	}
	return count
}

func GetArrangementsCount(boardSize int) int {
	return arrange(boardSize, []int{})
}
