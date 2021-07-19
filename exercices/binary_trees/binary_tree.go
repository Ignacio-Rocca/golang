package binary_trees

type Node struct {
	Val int
	Right *Node
	Left *Node
}
func NewNode(val int) *Node {
	return &Node{Val: val}
}

type BSTInterface interface {
	LookupOnBST(target int) bool
	InsertOnBST(val int) *Node
	Size() int
}

// Lookup is a func that given a BST, return true if a node with the target data is found in the tree. Recurs
// down the tree, chooses the left or right branch by comparing the target to each node.
func (root *Node) LookupOnBST(target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	} else if root.Val > target {
		// find in left
		return root.Left.LookupOnBST(target)
	} else {
		// find in right
		return root.Right.LookupOnBST(target)
	}
}

// Insert is a func that given a binary search tree and a number, insert a new node with the given number into the
// tree in the correct place
func (root *Node) InsertOnBST(val int) *Node {
	if root == nil {
		return NewNode(val)
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		if root.Left != nil {
			root.Left.InsertOnBST(val)
		} else {
			root.Left = NewNode(val)
		}
	} else {
		if root.Right != nil {
			root.Right.InsertOnBST(val)
		} else {
			root.Right = NewNode(val)
		}
	}

	return root
}

func (root *Node) Size() int {
	if root == nil {
		return 0
	}

	size := 1
	size += root.Left.Size()
	size += root.Right.Size()

	return size
}