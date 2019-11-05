package problem008

type node struct {
	Left  *node
	Right *node
	Value int
}

func NewNode(value int) *node {
	return &node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

func (tree *node) SetChildren(left *node, right *node) *node {
	tree.Left, tree.Right = left, right
	return tree
}

func (tree *node) isUnival() bool {
	if tree.Left == nil && tree.Right == nil {
		return true
	} else if tree.Right == nil {
		return tree.Left.isUnival() && tree.Left.Value == tree.Value
	} else if tree.Left == nil {
		return tree.Right.isUnival() && tree.Right.Value == tree.Value
	} else {
		subtreeChecks := make(chan bool, 2)
		go func() { subtreeChecks <- tree.Left.isUnival() }()
		go func() { subtreeChecks <- tree.Right.isUnival() }()
		return tree.Left.Value == tree.Right.Value && tree.Right.Value == tree.Value && <-subtreeChecks && <-subtreeChecks
	}
}

func (tree *node) CountUnival() int {
	univalCounts := make(chan int, 2)
	go func() {
		if tree.Left != nil {
			univalCounts <- tree.Left.CountUnival()
		} else {
			univalCounts <- 0
		}
	}()
	go func() {
		if tree.Right != nil {
			univalCounts <- tree.Right.CountUnival()
		} else {
			univalCounts <- 0
		}
	}()
	count := 0
	if tree.isUnival() {
		count += 1
	}
	return count + <-univalCounts + <-univalCounts
}
