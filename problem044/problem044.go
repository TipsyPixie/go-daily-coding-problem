package problem044

func mergeSort(values []int) (sortedValues []int, inversionCount int) {
	if len(values) <= 1 {
		sortedValues, inversionCount = values, 0
		return
	}
	mid := len(values) / 2
	sortedLefties, leftInversionCount := mergeSort(values[:mid])
	sortedRighties, rightInversionCount := mergeSort(values[mid:])
	inversionCount += leftInversionCount + rightInversionCount
	for leftIndex, rightIndex := 0, 0; leftIndex < len(sortedLefties) || rightIndex < len(sortedRighties); {
		switch {
		case rightIndex >= len(sortedRighties):
			sortedValues = append(sortedValues, sortedLefties[leftIndex])
			leftIndex++
		case leftIndex >= len(sortedLefties):
			sortedValues = append(sortedValues, sortedRighties[rightIndex])
			rightIndex++
		case sortedLefties[leftIndex] <= sortedRighties[rightIndex]:
			sortedValues = append(sortedValues, sortedLefties[leftIndex])
			leftIndex++
		case sortedLefties[leftIndex] > sortedRighties[rightIndex]:
			sortedValues = append(sortedValues, sortedRighties[rightIndex])
			inversionCount += len(sortedLefties) - leftIndex
			rightIndex++
		}
	}
	return
}

func CountInversions(values []int) (inversionCount int) {
	_, inversionCount = mergeSort(values)
	return
}
