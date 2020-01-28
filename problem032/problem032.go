package problem032

import "math"

type graph [][]float64

type edgeIterator func() (int, int, float64, bool)

func (thisGraph graph) iterateEdges() edgeIterator {
	row, column := 0, 0
	rowLessThan, columnLessThan := len(thisGraph), len(thisGraph[0])
	return func() (int, int, float64, bool) {
		from, to, weight, end := 0, 0, 0.0, true
		if row < rowLessThan && column < columnLessThan {
			from, to, weight, end = row, column, thisGraph[row][column], false
			column++
			if column == columnLessThan {
				column = 0
				row++
			}
		}
		return from, to, weight, end
	}
}

func (thisGraph graph) hasNegativeCycle() bool {
	shortestDistances := make([]float64, len(thisGraph), len(thisGraph))
	for i := 1; i < len(shortestDistances); i++ {
		// MaxFloat64 = unreachable
		shortestDistances[i] = math.MaxFloat64
	}

	for i := 0; i < len(thisGraph)-1; i++ {
		nextEdge := thisGraph.iterateEdges()
		for from, to, weight, end := nextEdge(); !end; from, to, weight, end = nextEdge() {
			if shortestDistances[from] < math.MaxFloat64 && shortestDistances[from]+weight < shortestDistances[to] {
				shortestDistances[to] = shortestDistances[from] + weight
			}
		}
	}

	nextEdge := thisGraph.iterateEdges()
	for from, to, weight, end := nextEdge(); !end; from, to, weight, end = nextEdge() {
		if shortestDistances[from] < math.MaxFloat64 && shortestDistances[from]+weight < shortestDistances[to] {
			return true
		}
	}
	return false
}

func HasArbitrage(rates [][]float64) bool {
	var negativeLogGraph graph
	negativeLogGraph = make([][]float64, len(rates), len(rates))
	for i := range negativeLogGraph {
		negativeLogGraph[i] = make([]float64, len(rates[0]), len(rates[0]))
	}

	for rowIndex, row := range rates {
		for columnIndex, cell := range row {
			negativeLogGraph[rowIndex][columnIndex] = -math.Log(cell)
		}
	}

	return negativeLogGraph.hasNegativeCycle()
}
