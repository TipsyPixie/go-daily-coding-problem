package problem012

func Run(maxStairCount int, stepSizes []int) int {
	waysToClimb := make([]int, maxStairCount+1, maxStairCount+1)
	waysToClimb[0] = 1
	for stairCount := 1; stairCount <= maxStairCount; stairCount++ {
		for _, stepSize := range stepSizes {
			if stepSize <= stairCount {
				waysToClimb[stairCount] += waysToClimb[stairCount-stepSize]
			}
		}
	}
	return waysToClimb[maxStairCount]
}
