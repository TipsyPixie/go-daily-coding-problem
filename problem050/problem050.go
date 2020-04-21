package problem050

import "errors"

type TreeNode interface {
	evaluate() (int, error)
}

type OperatorNode struct {
	value rune
	left  TreeNode
	right TreeNode
}

func (operatorNode OperatorNode) evaluate() (int, error) {
	leftValue, err := operatorNode.left.evaluate()
	if err != nil {
		return 0, err
	}
	rightValue, err := operatorNode.right.evaluate()
	if err != nil {
		return 0, err
	}
	switch operatorNode.value {
	case '+':
		return leftValue + rightValue, nil
	case '-':
		return leftValue - rightValue, nil
	case '*':
		return leftValue * rightValue, nil
	case '/':
		return leftValue / rightValue, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

type NumberNode struct {
	value int
}

func (numberNode NumberNode) evaluate() (int, error) {
	return numberNode.value, nil
}
