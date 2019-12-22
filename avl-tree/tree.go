package avl_tree

import "fmt"

type Tree struct {
	Top *Node
}

func (t *Tree) Add(v int) {
	var (
		newNode = &Node{Value: v}
	)
	if t.Top == nil {
		t.Top = newNode
	}
	node := t.Top

	// add
	for {
		if node.Value < v {
			if node.Left == nil {
				newNode.Parent = node
				node.Left = newNode
				break
			}
			node = node.Left
		} else if node.Value > v {
			if node.Right == nil {
				newNode.Parent = node
				node.Right = newNode
				break
			}
			node = node.Right
		}
		fmt.Printf("duplicate %d \n", v)
		break
	}
	// go up
	// update heights & rebalance
	for node != nil {
		if node.RightHeight() >= node.LeftHeight()+2 {
			rChild := node.Right
			// zig-zag
			if rChild.RightHeight() < rChild.LeftHeight() {
				t.RotateRight(rChild)
			}
			t.RotateLeft(node)
		}

	}

	// if balance == 2

	// RL

}

// RR -> rotate left
// Nodes: parent, unbalanced, r_child
// if parent = t.Top || if less or more => right/left
// parent = r_child
// unbalanced.Right = r_child.Left
// r_child.Left = unbalanced
func (t *Tree) RotateRight(n *Node) {

}

func (t *Tree) RotateLeft(node *Node) {
	var (
		parent     = node.Parent
		rightChild = node.Right
	)
	if parent == nil {
		t.Top = rightChild
	} else {
		if rightChild.Value < parent.Value {
			parent.Left = rightChild
		} else {
			parent.Right = rightChild
		}
	}
	node.Right = node.Right.Left
	rightChild.Left = node

	node.UpdHeight()
	rightChild.UpdHeight()
	parent.UpdHeight()
}

func (t *Tree) Delete(v int) {
	// top node case
	// not found
	// deletion from leave
	// deletion from the middle level with re-balance
}

func (t *Tree) Find(v int) {

}
