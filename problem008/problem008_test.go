package problem008

import "testing"

func TestBinaryTree_CountUnival(t *testing.T) {
	tree := NewNode(0).SetChildren(
		NewNode(1),
		NewNode(0).SetChildren(
			NewNode(1).SetChildren(
				NewNode(1),
				NewNode(1),
			),
			NewNode(0),
		),
	)
	if tree.CountUnival() != 5 {
		t.FailNow()
	}
}

func TestBinaryTree_CountUnival2(t *testing.T) {
	tree := NewNode(5).SetChildren(
		NewNode(1).SetChildren(
			NewNode(5),
			NewNode(5),
		),
		NewNode(5).SetChildren(
			nil,
			NewNode(5),
		),
	)
	if tree.CountUnival() != 4 {
		t.FailNow()
	}
}
