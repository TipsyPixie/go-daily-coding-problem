package problem058

func Search(numbers []int, targetNumber int) *int {
	return modifiedBinarySearch(numbers, targetNumber, 0)
}

func getIntPointer(value int) *int {
	return &value
}

func modifiedBinarySearch(numbers []int, targetNumber int, outerIndex int) *int {
	if len(numbers) == 0 {
		return nil
	}

	midIndex := (len(numbers) - 1) / 2
	if targetNumber == numbers[midIndex] {
		return getIntPointer(outerIndex + midIndex)
	}

	firstNumber, midNumber, lastNumber := numbers[0], numbers[midIndex], numbers[len(numbers)-1]
	if firstNumber <= midNumber {
		if firstNumber <= targetNumber && targetNumber < midNumber {
			return modifiedBinarySearch(numbers[:midIndex], targetNumber, 0)
		}
		return modifiedBinarySearch(numbers[midIndex+1:], targetNumber, midIndex+1)
	} else {
		if midNumber < targetNumber && targetNumber <= lastNumber {
			return modifiedBinarySearch(numbers[midIndex+1:], targetNumber, midIndex+1)
		}
		return modifiedBinarySearch(numbers[:midIndex], targetNumber, 0)
	}
}
