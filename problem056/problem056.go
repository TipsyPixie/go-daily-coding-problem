package problem056

const NO_COLOR = -1

func isValid(nodeId int, color int, nodeColors []int, adjacency [][]bool) bool {
	for otherNodeId, isAdjacent := range adjacency[nodeId] {
		if isAdjacent && color == nodeColors[otherNodeId] {
			return false
		}
	}
	return true
}

func resolve(nodeId int, nodeColors []int, adjacency [][]bool, maxColor int) bool {
	if nodeId >= len(nodeColors) {
		return true
	}
	for color := 0; color < maxColor; color++ {
		if isValid(nodeId, color, nodeColors, adjacency) {
			nodeColors[nodeId] = color
			if resolve(nodeId+1, nodeColors, adjacency, maxColor) {
				return true
			}
			nodeColors[nodeId] = NO_COLOR
		}
	}
	return false
}

func HasSolution(adjacency [][]bool, maxColor int) bool {
	nodeColors := make([]int, maxColor, maxColor)
	for i := range nodeColors {
		nodeColors[i] = NO_COLOR
	}
	return resolve(0, nodeColors, adjacency, maxColor)
}
