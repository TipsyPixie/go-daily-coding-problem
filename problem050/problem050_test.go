package problem050

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tree := OperatorNode{
		value: '*',
		left: OperatorNode{
			value: '+',
			left:  TreeNode(&NumberNode{3}),
			right: TreeNode(&NumberNode{2}),
		},
		right: OperatorNode{
			value: '+',
			left:  TreeNode(&NumberNode{4}),
			right: TreeNode(&NumberNode{5}),
		},
	}
	result, err := tree.evaluate()
	correctResult := (3 + 2) * (4 + 5)
	if err != nil || result != correctResult {
		t.Error(err)
		t.Error(fmt.Sprintf("%d != %d", result, correctResult))
		t.Fail()
	}
}
