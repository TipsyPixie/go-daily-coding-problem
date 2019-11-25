package problem026

type linkedList struct {
	value int
	next *linkedList
}

func (node *linkedList) FromLast(reverseIndex int) *linkedList {
	formerCursor, latterCursor := node, node
	for i := 0; i < reverseIndex; i++ {
		latterCursor = latterCursor.next
	}
	for latterCursor.next != nil {
		formerCursor, latterCursor = formerCursor.next, latterCursor.next
	}
	return formerCursor
}
