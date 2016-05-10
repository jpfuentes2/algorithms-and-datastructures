package bstree

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
type Visitor func(*Node)

// Node is a binary search tree node
type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

// Tree is a binary search tree
type Tree struct {
	Root *Node
	Size int
}

// New creates a new BinarySearchTree
func New() *Tree {
	return &Tree{}
}

// Insert inserts the key into this sub-tree. error if the key is a duplicate
func (n *Node) Insert(key int) error {
	if key == n.Key {
		// do not allow duplicates
		return ErrDuplicateKey
	} else if key < n.Key {
		// if key is < n.key then insert to left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	} else {
		// if key is > n.key then insert to right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	}

	return nil
}

// IsLeaf returns true if this node has no children
func (n *Node) IsLeaf() bool {
	n.Left == nil && n.Right == nil
}

// IsEmpty returns true if the root is nil
func (t *Tree) IsEmpty() bool {
	return t.Len() == 0
}

// Len is the number of nodes in this tree
func (t *Tree) Len() int {
	return t.Size
}

// PreOrder walks the tree pre-order
func (t *Tree) PreOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *Node) {
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
func (t *Tree) InOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *Node) {
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
func (t *Tree) PostOrder(fn Visitor) {
	if t.IsEmpty() {
		return
	}

	var recurse Visitor
	recurse = func(n *Node) {
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
func (t *Tree) Min() *Node {
	if t.IsEmpty() {
		return nil
	}

	var recurse func(node *Node) *Node
	recurse = func(node *Node) *Node {
		if node.Left == nil {
			return node
		}

		return recurse(node.Left)
	}

	return recurse(t.Root)
}

// Max gives the maximum key of this tree
func (t *Tree) Max() *Node {
	if t.IsEmpty() {
		return nil
	}

	var recurse func(node *Node) *Node
	recurse = func(node *Node) *Node {
		if node.Right == nil {
			return node
		}

		return recurse(node.Right)
	}

	return recurse(t.Root)
}

// Insert adds the key as a node. error if tried to add a duplicate key
func (t *Tree) Insert(key int) (err error) {
	if t.IsEmpty() {
		t.Root = &Node{Key: key}
		t.Size++
		return
	}

	if err = t.Root.Insert(key); err == nil {
		t.Size++
	}
	return
}

// Contains tests for existence of the search key
func (t *Tree) Contains(key int) bool {
	_, err := t.Get(key)
	return err == nil
}

// Get gets the node with the given key. error if not found
func (t *Tree) Get(key int) (*Node, error) {
	var recurse func(key int, node *Node) *Node
	recurse = func(key int, node *Node) *Node {
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

	if node := recurse(key, t.Root); node != nil {
		return node, nil
	}

	return nil, ErrNodeNotFound
}

// Height gives the height of this tree
func (t *Tree) Height() int {
	if t.IsEmpty() {
		return 0
	}

	var recurse func(n *Node) int
	recurse = func(n *Node) int {
		if n == nil {
			return -1
		}

		return util.Max(recurse(n.Left), recurse(n.Right)) + 1
	}

	return recurse(t.Root)
}

// Remove removes the key from the tree. error if not found or tree is empty
func (t *Tree) Remove(key int) error {
	if t.IsEmpty() {
		return ErrNodeNotFound
	}

	var recurse func(key int, node *Node) *Node
	recurse = func(key int, node *Node) *Node {
		if node == nil {
			return nil
		}

		if key < node.Key {
			node.Left = recurse(key, node.Left)
		} else if key > node.Key {
			node.Right = recurse(key, node.Right)
		} else {
			if node.Left == nil {
				return node.Right
			}
			if node.Right == nil {
				return node.Left
			}

			t := node
			node = min(t.Right)
			node.Right = deleteMin(t.Right)
		}

		return node
	}

	if node := recurse(key, t.Root); node != nil {
		t.Size--
		return
	}

	return ErrNodeNotFound
}
