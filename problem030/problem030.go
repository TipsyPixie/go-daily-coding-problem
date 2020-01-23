package problem030

func MeasureTrappedWater(wallHeights []int) int {
	trappedWater := 0
	leftCursor, rightCursor := 0, len(wallHeights)-1
	leftMaxHeight, rightMaxHeight := 0, 0
	for leftCursor <= rightCursor {
		if leftHeight, rightHeight := wallHeights[leftCursor], wallHeights[rightCursor]; leftHeight < rightHeight {
			if leftHeight > leftMaxHeight {
				leftMaxHeight = leftHeight
			} else {
				trappedWater += leftMaxHeight - leftHeight
			}
			leftCursor++
		} else {
			if rightHeight > rightMaxHeight {
				rightMaxHeight = rightHeight
			} else {
				trappedWater += rightMaxHeight - rightHeight
			}
			rightCursor--
		}
	}

	return trappedWater
}
