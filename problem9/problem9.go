package problem9

func Run(input []int) int {
	maxIncludingLast, maxExcludingLast := 0, 0
	for _, n := range input {
		oldMaxIncludingLast, oldMaxExcludingLast := maxIncludingLast, maxExcludingLast
		maxIncludingLast = oldMaxExcludingLast + n
		if oldMaxIncludingLast > oldMaxExcludingLast {
			maxExcludingLast = oldMaxIncludingLast
		}
	}
	if maxIncludingLast > maxExcludingLast {
		return maxIncludingLast
	} else {
		return maxExcludingLast
	}
}
