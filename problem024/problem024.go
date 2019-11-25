package problem024

type binaryTree struct {
	lockedDescendantsCount int
	locked                 bool
	parent                 *binaryTree
	left                   *binaryTree
	right                  *binaryTree
}

func NewBinaryTree(left *binaryTree, right *binaryTree) *binaryTree {
	node := binaryTree{
		parent:                 nil,
		lockedDescendantsCount: 0,
		locked:                 false,
		left:                   left,
		right:                  right,
	}
	if right != nil {
		right.parent = &node
	}
	if left != nil {
		left.parent = &node
	}
	return &node
}

func (node *binaryTree) IsLocked() bool {
	return node.locked
}

func (node *binaryTree) Lock() bool {
	if node.lockedDescendantsCount > 0 {
		return false
	}
	ascendants := make([]*binaryTree, 0, 64)
	for cursor := node.parent; cursor != nil; cursor = cursor.parent {
		if cursor.locked {
			return false
		}
		ascendants = append(ascendants, cursor)
	}
	for _, ascendant := range ascendants {
		ascendant.lockedDescendantsCount++
	}
	node.locked = true
	return true
}

func (node *binaryTree) Unlock() bool {
	if node.lockedDescendantsCount > 0 {
		return false
	}
	ascendants := make([]*binaryTree, 0, 64)
	for cursor := node.parent; cursor != nil; cursor = cursor.parent {
		if cursor.locked {
			return false
		}
		ascendants = append(ascendants, cursor)
	}
	for _, ascendant := range ascendants {
		ascendant.lockedDescendantsCount--
	}
	node.locked = false
	return true
}
