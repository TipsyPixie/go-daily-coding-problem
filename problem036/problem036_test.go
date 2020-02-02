package problem036

import "testing"

func TestBinarySearchTree_FindSecondLargest(t *testing.T) {
	tree := NewBstNode(10).Insert(7).Insert(5).Insert(14).Insert(9)
	result := tree.FindSecondLargest().value
	if result != 10 {
		t.Fail()
		t.Log(result)
	}

	tree = NewBstNode(1).Insert(2).Insert(5).Insert(4).Insert(3).Insert(11)
	result = tree.FindSecondLargest().value
	if result != 5 {
		t.Fail()
		t.Log(result)
	}
}
