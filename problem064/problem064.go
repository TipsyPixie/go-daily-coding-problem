package problem064

type coordinates struct {
	x int
	y int
}

func CountKnightsPaths(boardSize int) int {
	totalCountChannel := make(chan int, boardSize*boardSize)
	countPaths := func(x int, y int) {
		visitedCoordinates := make(map[coordinates]bool)
		initialCoordinates := coordinates{x: x, y: y}
		visitedCoordinates[initialCoordinates] = true
		totalCountChannel <- countPathsByBruteforce(visitedCoordinates, initialCoordinates, boardSize)
	}

	for initialX := 0; initialX < boardSize; initialX++ {
		for initialY := 0; initialY < boardSize; initialY++ {
			go countPaths(initialX, initialY)
		}
	}

	totalCount := 0
	for i := 0; i < boardSize*boardSize; i++ {
		select {
		case partialCount := <-totalCountChannel:
			totalCount += partialCount
		}
	}
	return totalCount
}

func getAllNextCoordinates(currentCoordinates coordinates) []coordinates {
	return []coordinates{
		{x: currentCoordinates.x - 2, y: currentCoordinates.y - 1},
		{x: currentCoordinates.x - 1, y: currentCoordinates.y - 2},
		{x: currentCoordinates.x + 1, y: currentCoordinates.y - 2},
		{x: currentCoordinates.x + 2, y: currentCoordinates.y - 1},
		{x: currentCoordinates.x + 2, y: currentCoordinates.y + 1},
		{x: currentCoordinates.x + 1, y: currentCoordinates.y + 2},
		{x: currentCoordinates.x - 1, y: currentCoordinates.y + 2},
		{x: currentCoordinates.x - 2, y: currentCoordinates.y + 1},
	}
}

func countPathsByBruteforce(visitedCoordinates map[coordinates]bool, currentCoordinates coordinates, boardSize int) int {
	if len(visitedCoordinates) == boardSize*boardSize {
		return 1
	}

	pathCount := 0
	allNextCoordinates := getAllNextCoordinates(currentCoordinates)
	for _, nextCoordinates := range allNextCoordinates {
		if nextCoordinates.x >= 0 && nextCoordinates.x <= boardSize-1 &&
			nextCoordinates.y >= 0 && nextCoordinates.y <= boardSize-1 &&
			!visitedCoordinates[nextCoordinates] {
			visitedCoordinates[nextCoordinates] = true
			if nextPathCount := countPathsByBruteforce(visitedCoordinates, nextCoordinates, boardSize); nextPathCount > 0 {
				pathCount += nextPathCount
			}
			delete(visitedCoordinates, nextCoordinates)
		}
	}
	return pathCount
}
