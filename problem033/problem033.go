package problem033

func insertSorted(list []int, value int) []int {
	i := 0
	for i < len(list) && (list)[i] < value {
		i++
	}
	list = append(list, 0)
	copy(list[i+1:], list[i:])
	list[i] = value
	return list
}

func GetMedians(list []int) []float32 {
	medians, sortedList := make([]float32, len(list), len(list)), make([]int, 0, 0)
	for i, value := range list {
		sortedList = insertSorted(sortedList, value)
		if len(sortedList)%2 == 1 {
			medians[i] = float32(sortedList[len(sortedList)/2])
		} else if len(sortedList) == 0 {
			medians[i] = float32(value)
		} else {
			medians[i] = float32(sortedList[len(sortedList)/2-1]+sortedList[len(sortedList)/2]) / 2
		}
	}
	return medians
}
