package avl_tree

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
	Height int
}

func (n *Node) UpdHeight() {
	rightHeight := n.RightHeight()
	leftHeight := n.LeftHeight()
	max := func() int {
		if leftHeight > rightHeight {
			return leftHeight
		}
		return rightHeight
	}()
	n.Height = max + 1
}

func (n *Node) LeftHeight() int {
	if n.Left == nil {
		return -1
	}
	return n.Left.Height
}

func (n *Node) RightHeight() int {
	if n.Right == nil {
		return -1
	}
	return n.Right.Height
}
