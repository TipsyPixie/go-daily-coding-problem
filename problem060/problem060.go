package problem060

func CanBePartitioned(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	total := 0
	for _, number := range numbers {
		total += number
	}
	if total%2 > 0 {
		return false
	}
	return canBePartitioned(numbers, 0, total/2)
}

func canBePartitioned(numbers []int, firstSum int, halfTotal int) bool {
	for i, number := range numbers {
		tempFirstSum := firstSum + number
		switch {
		case tempFirstSum == halfTotal:
			return true
		case tempFirstSum < halfTotal:
			if canBePartitioned(append(numbers[:i], numbers[i:]...), tempFirstSum, halfTotal) {
				return true
			}
		}
	}
	return false
}
