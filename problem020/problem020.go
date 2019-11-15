package problem20

type linkedList struct {
	next  *linkedList
	value int
}

func ToLinkedList(values []int) *linkedList {
	var lastOne *linkedList
	var firstOne *linkedList
	for _, value := range values {
		newLinkedList := &linkedList{
			next:  nil,
			value: value,
		}
		if lastOne == nil {
			firstOne = newLinkedList
			lastOne = firstOne
		} else {
			lastOne.next = newLinkedList
			lastOne = lastOne.next
		}
	}
	return firstOne
}

func Run(firstList *linkedList, secondList *linkedList) int {
	firstLength, secondLength := 1, 1
	for cursor := firstList; cursor.next != nil; cursor = cursor.next {
		firstLength++
	}
	for cursor := secondList; cursor.next != nil; cursor = cursor.next {
		secondLength++
	}

	lengthDifference := 0
	var shorterOne, longerOne *linkedList
	if firstLength >= secondLength {
		lengthDifference = firstLength - secondLength
		shorterOne, longerOne = secondList, firstList
	} else {
		lengthDifference = secondLength - firstLength
		shorterOne, longerOne = firstList, secondList
	}

	for i := 0; i < lengthDifference; i++ {
		longerOne = longerOne.next
	}
	for shorterOne.value != longerOne.value {
		shorterOne = shorterOne.next
		longerOne = longerOne.next
	}
	return shorterOne.value
}
