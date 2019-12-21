package avl_tree

type Node struct {
	Value   int
	Left    *Node
	Right   *Node
	Balance int
}

type Tree struct {
	Top *Node
	// Sentinel
}

func (t *Tree) Add(v int) {
	var (
		unbalanced *Node
		newNode    = &Node{Value: v}
	)
	if t.Top == nil {
		t.Top = newNode
	}
	node := t.Top
	for {
		if node.Value < v {
			if node.Left == nil {
				node.Left = newNode
				break
			}
			node = node.Left
		} else if node.Value > v {
			if node.Right == nil {
				node.Right = newNode
				break
			}
			node = node.Right
		}
		break
	}
	if unbalanced != nil {
		rotate
	}
}

func (t *Tree) Delete(v int) {
	// top node case
	// not found
	// deletion from leave
	// deletion from the middle level with re-balance
}

func (t *Tree) Find(v int) {

}
