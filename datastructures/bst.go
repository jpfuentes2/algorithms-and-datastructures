package datastructures

import (
	"errors"

	"github.com/jpfuentes2/algorithms-and-datastructures/util"
)

var (
	// ErrDuplicateKey is an error when a duplicate key is inserted
	ErrDuplicateKey = errors.New("Cannot insert duplicate key")

	// ErrNodeNotFound is an error when the node is not found in the tree
	ErrNodeNotFound = errors.New("Cannot find node in tree")
)

// Visitor is a type alias for an anon fn receiving a node
type Visitor func(*BSTreeNode)

// BSTreeNode is a binary search tree node
type BSTreeNode struct {
	Left  *BSTreeNode
	Key   int
	Right *BSTreeNode
}

// BSTree is a binary search tree
type BSTree struct {
	Root *BSTreeNode
	Size int
}

// NewBSTree creates a new BinarySearchTree
func NewBSTree() *BSTree {
	return &BSTree{}
}

// Insert inserts the key into this sub-tree. error if the key is a duplicate
func (n *BSTreeNode) Insert(key int) error {
	if key == n.Key {
		// do not allow duplicates
		return ErrDuplicateKey
	} else if key < n.Key {
		// if key is < n.key then insert to left
		if n.Left == nil {
			n.Left = &BSTreeNode{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else {
		// if key is > n.key then insert to right
		if n.Right == nil {
			n.Right = &BSTreeNode{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}

	return nil
}

// IsEmpty returns true if the root is nil
func (t *BSTree) IsEmpty() bool {
	return t.Len() == 0
}

// Len is the number of nodes in this tree
func (t *BSTree) Len() int {
	return t.Size
}

// PreOrder walks the tree pre-order
func (t *BSTree) PreOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *BSTreeNode) {
		if n == nil {
			return
		}
		fn(n)
		recurse(n.Left)
		recurse(n.Right)
	}

	recurse(t.Root)
}

// InOrder walks the tree in-order
func (t *BSTree) InOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *BSTreeNode) {
		if n == nil {
			return
		}
		recurse(n.Left)
		fn(n)
		recurse(n.Right)
	}

	recurse(t.Root)
}

// PostOrder walks the tree in post-order
func (t *BSTree) PostOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *BSTreeNode) {
		if n == nil {
			return
		}
		recurse(n.Left)
		recurse(n.Right)
		fn(n)
	}

	recurse(t.Root)
}

// Min gives the minimum key of this tree
func (t *BSTree) Min() *BSTreeNode {
	if t.IsEmpty() {
		return nil
	}

	var recurse func(node *BSTreeNode) *BSTreeNode
	recurse = func(node *BSTreeNode) *BSTreeNode {
		if node.Left == nil {
			return node
		}

		return recurse(node.Left)
	}

	return recurse(t.Root)
}

// Max gives the maximum key of this tree
func (t *BSTree) Max() *BSTreeNode {
	if t.IsEmpty() {
		return nil
	}

	var recurse func(node *BSTreeNode) *BSTreeNode
	recurse = func(node *BSTreeNode) *BSTreeNode {
		if node.Right == nil {
			return node
		}

		return recurse(node.Right)
	}

	return recurse(t.Root)
}

// Insert adds the key as a node. error if tried to add a duplicate key
func (t *BSTree) Insert(key int) (err error) {
	if t.IsEmpty() {
		t.Root = &BSTreeNode{Key: key}
		t.Size++
		return
	}

	if err = t.Root.Insert(key); err == nil {
		t.Size++
	}
	return
}

// Contains tests for existence of the search key
func (t *BSTree) Contains(key int) bool {
	var recurse func(key int, node *BSTreeNode) *BSTreeNode
	recurse = func(key int, node *BSTreeNode) *BSTreeNode {
		if node == nil {
			return nil
		}

		if key < node.Key {
			return recurse(key, node.Left)
		} else if key > node.Key {
			return recurse(key, node.Right)
		} else {
			return node
		}
	}

	return recurse(key, t.Root) != nil
}

// Height gives the height of this tree
func (t *BSTree) Height() int {
	if t.IsEmpty() {
		return 0
	}

	var recurse func(n *BSTreeNode) int
	recurse = func(n *BSTreeNode) int {
		if n == nil {
			return -1
		}

		return util.Max(recurse(n.Left), recurse(n.Right)) + 1
	}

	return recurse(t.Root)
}

// Remove removes the node from the tree
func (t *BSTree) Remove(n *BSTreeNode) error {
	return ErrNodeNotFound
}
