package problem026

import "testing"

func TestLinkedList_FromLast(t *testing.T) {
	node := &linkedList{
		value: 3,
		next: &linkedList{
			value: 2,
			next: &linkedList{
				value: 1,
				next: &linkedList{
					value: 0,
					next:  nil,
				},
			},
		},
	}
	if node.FromLast(0).value != 0 {
		t.Log(node.FromLast(0))
		t.FailNow()
	}
	if node.FromLast(2).value != 2 {
		t.Log(node.FromLast(2))
		t.FailNow()
	}
}
