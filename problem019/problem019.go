package problem019

import (
	"math"
)

func minInt(values ...int) int {
	minimum := math.MaxInt32
	for _, value := range values {
		if minimum > value {
			minimum = value
		}
	}
	return minimum
}

func Run(costMap [][]int) int {
	if len(costMap) == 0 || len(costMap[0]) == 0 {
		return 0
	}
	colorCount := len(costMap[0])
	lastMinCosts, minCosts := make([]int, colorCount, colorCount), make([]int, colorCount, colorCount)
	for colorIndex, cost := range costMap[0] {
		lastMinCosts[colorIndex] = cost
	}
	for houseIndex := 1; houseIndex < len(costMap); houseIndex++ {
		for colorIndex, cost := range costMap[houseIndex] {
			minCost := math.MaxInt32
			for lastColorIndex, lastMinCost := range lastMinCosts {
				if lastColorIndex != colorIndex && minCost > lastMinCost+cost {
					minCost = lastMinCost + cost
				}
			}
			minCosts[colorIndex] = minCost
		}
		lastMinCosts, minCosts = minCosts, lastMinCosts
	}
	return minInt(lastMinCosts...)
}
