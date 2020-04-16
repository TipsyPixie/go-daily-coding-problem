package problem042

func GetSubset(values []int, sum int) []int {
	if len(values) == 0 {
		if sum == 0 {
			return []int{}
		} else {
			return nil
		}
	}
	if subsetWithoutFirstItem := GetSubset(values[1:], sum); subsetWithoutFirstItem != nil {
		return subsetWithoutFirstItem
	}
	if subsetWithFirstItem := GetSubset(values[1:], sum-values[0]); subsetWithFirstItem != nil {
		return append(values[0:1], subsetWithFirstItem...)
	}
	return nil
}
