package problem024

import (
	"os"
	"testing"
)

var tree *binaryTree

func TestMain(m *testing.M) {
	tree = NewBinaryTree(
		NewBinaryTree(nil, nil),
		NewBinaryTree(
			NewBinaryTree(nil, nil),
			NewBinaryTree(nil, nil),
		),
	)
	os.Exit(m.Run())
}

func TestBinaryTree_Lock(t *testing.T) {
	if tree.right.IsLocked() {
		t.Log("right node already locked")
		t.FailNow()
	}
	if !tree.right.Lock() || !tree.right.IsLocked() || tree.lockedDescendantsCount != 1 {
		t.Log("failed to lock right node")
		t.FailNow()
	}
	if tree.Lock() || tree.right.right.Lock() {
		t.Log("invalid lock")
		t.FailNow()
	}
}

func TestBinaryTree_Unlock(t *testing.T) {
	if tree.Unlock() || tree.right.right.Unlock() {
		t.Log("invalid lock")
		t.FailNow()
	}
	if !tree.right.Unlock() {
		t.Log("failed to unlock right node")
		t.FailNow()
	}
}
