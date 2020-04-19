package problem048

type TreeNode struct {
	value      rune
	leftChild  *TreeNode
	rightChild *TreeNode
}

func (thisNode *TreeNode) Equal(other *TreeNode) bool {
	if thisNode == nil || other == nil {
		return thisNode == nil && other == nil
	}
	childEqualityChannel := make(chan bool, 2)
	go func() {
		childEqualityChannel <- thisNode.leftChild.Equal(other.leftChild)
	}()
	childEqualityChannel <- thisNode.rightChild.Equal(other.rightChild)

	signalCount := 0
	for signalCount < 2 {
		select {
		case childEquality := <-childEqualityChannel:
			signalCount++
			if !childEquality {
				return false
			}
		}
	}
	return thisNode.value == other.value
}

type Order int

const (
	PREORDER  = Order(0)
	INORDER   = Order(1)
	POSTORDER = Order(2)
)

func FromValues(values []rune, order Order) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	var leftChildValues, rightChildValues []rune
	var rootValue rune

	// presume the tree is balanced
	switch order {
	case PREORDER:
		rootValue = values[0]
		childValues := values[1:]
		mid := len(childValues) / 2
		leftChildValues, rightChildValues = childValues[:mid], childValues[mid:]
	case INORDER:
		mid := len(values) / 2
		rootValue = values[mid]
		leftChildValues, rightChildValues = values[:mid], values[mid+1:]
	case POSTORDER:
		rootValue = values[len(values)-1]
		childValues := values[:len(values)-1]
		mid := len(childValues) / 2
		leftChildValues, rightChildValues = childValues[:mid], childValues[mid:]
	}

	return &TreeNode{
		value:      rootValue,
		leftChild:  FromValues(leftChildValues, order),
		rightChild: FromValues(rightChildValues, order),
	}
}
