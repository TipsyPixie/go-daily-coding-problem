package problem036

type binarySearchTree struct {
	value      int
	leftChild  *binarySearchTree
	rightChild *binarySearchTree
}

func NewBstNode(value int) *binarySearchTree {
	return &binarySearchTree{
		value:      value,
		leftChild:  nil,
		rightChild: nil,
	}
}

func (root *binarySearchTree) Insert(value int) *binarySearchTree {
	tree := root
	for {
		if value > tree.value {
			if tree.rightChild == nil {
				tree.rightChild = NewBstNode(value)
				return root
			}
			tree = tree.rightChild
		} else {
			if tree.leftChild == nil {
				tree.leftChild = NewBstNode(value)
				return root
			}
			tree = tree.leftChild
		}
	}
}

func (root *binarySearchTree) FindSecondLargest() *binarySearchTree {
	var parentTree *binarySearchTree
	tree := root
	for tree.rightChild != nil {
		parentTree = tree
		tree = tree.rightChild
	}
	return parentTree
}
