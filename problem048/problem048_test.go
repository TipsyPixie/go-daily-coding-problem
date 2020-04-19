package problem048

import (
	"fmt"
	"testing"
)

func TestFromValues(t *testing.T) {
	testArgs := [][]rune{
		{'a', 'b', 'd', 'e', 'c', 'f', 'g'},
		{'d', 'b', 'e', 'a', 'f', 'c', 'g'},
	}
	testArgs2 := []Order{PREORDER, INORDER}
	correctAnswer := TreeNode{
		value: 'a',
		leftChild: &TreeNode{
			value: 'b',
			leftChild: &TreeNode{
				value:      'd',
				leftChild:  nil,
				rightChild: nil,
			},
			rightChild: &TreeNode{
				value:      'e',
				leftChild:  nil,
				rightChild: nil,
			},
		},
		rightChild: &TreeNode{
			value: 'c',
			leftChild: &TreeNode{
				value:      'f',
				leftChild:  nil,
				rightChild: nil,
			},
			rightChild: &TreeNode{
				value:      'g',
				leftChild:  nil,
				rightChild: nil,
			},
		},
	}
	for i, _ := range testArgs {
		answer := FromValues(testArgs[i], testArgs2[i])
		if !answer.Equal(&correctAnswer) {
			t.Log(fmt.Sprintf("failed at %v, %v", testArgs[i], testArgs2[i]))
			t.Fail()
		}
	}
}
