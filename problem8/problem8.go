package problem8

type binaryTree struct {
    Left  *binaryTree
    Right *binaryTree
    Value int
}

func NewNode(value int) *binaryTree {
    return &binaryTree{
        Left:  nil,
        Right: nil,
        Value: value,
    }
}

func (tree *binaryTree) SetChildren(left *binaryTree, right *binaryTree) *binaryTree {
    tree.Left = left
    tree.Right = right
    return tree
}

func (tree *binaryTree) isUnival() (bool, int) {
    if tree.Left == nil && tree.Right == nil {
        return true, tree.Value
    } else if tree.Right == nil {
        isLeftUnival, leftValue := tree.Left.isUnival()
        return isLeftUnival && leftValue == tree.Value, tree.Value
    } else if tree.Left == nil {
        isRightUnival, rightValue := tree.Right.isUnival()
        return isRightUnival && rightValue == tree.Value, tree.Value
    } else {
        ch1, ch2 := make(chan bool, 2), make(chan int, 2)
        go func() {
            isUnival, uniValue := tree.Left.isUnival()
            ch1 <- isUnival
            ch2 <- uniValue
        }()
        go func() {
            isUnival, uniValue := tree.Right.isUnival()
            ch1 <- isUnival
            ch2 <- uniValue
        }()
        isSubtreeUnival1, isSubtreeUnival2, subtreeUnival1, subtreeUnival2 := <-ch1, <-ch1, <-ch2, <-ch2
        return isSubtreeUnival1 && isSubtreeUnival2 && subtreeUnival1 == subtreeUnival2 && subtreeUnival2 == tree.Value, tree.Value
    }
}

func (tree *binaryTree) CountUnival() int {
    leftUnivalCount, rightUnivalCount := make(chan int, 1), make(chan int, 1)
    go func() {
        if tree.Left != nil {
            leftUnivalCount <- tree.Left.CountUnival()
            return
        }
        close(leftUnivalCount)
    }()
    go func() {
        if tree.Right != nil {
            rightUnivalCount <- tree.Right.CountUnival()
            return
        }
        close(rightUnivalCount)
    }()
    count := 0
    if isUnival, _ := tree.isUnival(); isUnival {
        count += 1
    }
    return count + <-leftUnivalCount + <-rightUnivalCount
}
