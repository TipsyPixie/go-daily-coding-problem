package problem20

import "testing"

func TestRun(t *testing.T) {
	list1, list2 := ToLinkedList([]int{3, 7, 8, 10}), ToLinkedList([]int{14, 0, 5, 99, 1, 8, 10})
	if Run(list1, list2) != 8 {
		t.FailNow()
	}
}
