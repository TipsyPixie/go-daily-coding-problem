package problem049

func MaximizeContiguousSum(values []int) int {
	// ex) [10, 20, -50, 24, 1, -10, 20, -6] => [10, 20, -50] [24, 1, -10, 20] [-6]
	totalSum := 0
	for _, value := range values {
		totalSum += value
	}

	leftExcludedMinSum, leftExcludedSum, leftExcludedEndsAt := 0, 0, 0
	rightExcludedMinSum, rightExcludedSum, rightExcludedStartsAt := totalSum, totalSum, 0
	for i := range values {
		if leftExcludedSum += values[i]; leftExcludedSum < leftExcludedMinSum {
			leftExcludedEndsAt = i + 1
		}
		if rightExcludedSum -= values[i]; rightExcludedSum < rightExcludedMinSum {
			rightExcludedStartsAt = i + 1
		}
	}

	if leftExcludedEndsAt >= rightExcludedStartsAt {
		return 0
	}
	maxContiguousSum := 0
	for _, value := range values[leftExcludedEndsAt:rightExcludedStartsAt] {
		maxContiguousSum += value
	}
	return maxContiguousSum
}
