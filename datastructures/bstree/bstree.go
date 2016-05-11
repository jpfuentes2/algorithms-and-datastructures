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
	left  *Node
	Key   int
	right *Node
}

// Tree is a binary search tree
type Tree struct {
	root *Node
	size int
}

// Insert inserts the key into this sub-tree. error if the key is a duplicate
func (n *Node) Insert(key int) error {
	if key == n.Key {
		// do not allow duplicates
		return ErrDuplicateKey
	} else if key < n.Key {
		// if key is < n.key then insert to left
		if n.left == nil {
			n.left = &Node{Key: key}
		} else {
			n.left.Insert(key)
		}
	} else {
		// if key is > n.key then insert to right
		if n.right == nil {
			n.right = &Node{Key: key}
		} else {
			n.right.Insert(key)
		}
	}

	return nil
}

// IsLeaf returns true if this node has no children
func (n *Node) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

// New creates a new BinarySearchTree
func New() *Tree {
	return &Tree{}
}

// IsEmpty returns true if the root is nil
func (t *Tree) IsEmpty() bool {
	return t.Len() == 0
}

// Len is the number of nodes in this tree
func (t *Tree) Len() int {
	return t.size
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
		recurse(n.left)
		recurse(n.right)
	}

	recurse(t.root)
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
		recurse(n.left)
		fn(n)
		recurse(n.right)
	}

	recurse(t.root)
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
		recurse(n.left)
		recurse(n.right)
		fn(n)
	}

	recurse(t.root)
}

// Min gives the minimum key of this tree
func (t *Tree) Min() *Node {
	if t.IsEmpty() {
		return nil
	}

	return t.min(t.root)
}

func (t *Tree) min(node *Node) *Node {
	if node.left == nil {
		return node
	}

	return t.min(node.left)
}

// Max gives the maximum key of this tree
func (t *Tree) Max() *Node {
	if t.IsEmpty() {
		return nil
	}

	var recurse func(node *Node) *Node
	recurse = func(node *Node) *Node {
		if node.right == nil {
			return node
		}

		return recurse(node.right)
	}

	return recurse(t.root)
}

// Insert adds the key as a node. error if tried to add a duplicate key
func (t *Tree) Insert(key int) (err error) {
	if t.IsEmpty() {
		t.root = &Node{Key: key}
		t.size++
		return
	}

	if err = t.root.Insert(key); err == nil {
		t.size++
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
			return recurse(key, node.left)
		} else if key > node.Key {
			return recurse(key, node.right)
		} else {
			return node
		}
	}

	if node := recurse(key, t.root); node != nil {
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

		return util.Max(recurse(n.left), recurse(n.right)) + 1
	}

	return recurse(t.root)
}

func (t *Tree) deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = t.deleteMin(node.left)
	return node
}

// Remove removes the key from the tree. error if not found or tree is empty
func (t *Tree) Remove(key int) (*Node, error) {
	if t.IsEmpty() {
		return nil, ErrNodeNotFound
	}

	if node := t.remove(key, t.root); node != nil {
		t.size--
		return node, nil
	}

	return nil, ErrNodeNotFound
}

func (t *Tree) remove(key int, node *Node) *Node {
	if node == nil {
		return nil
	}

	if key < node.Key {
		removed := t.remove(key, node.left)
		node.left = removed
		node = removed
	} else if key > node.Key {
		removed := t.remove(key, node.right)
		node.right = removed
		node = removed
	} else {
		// remove a leaf node straight-up
		if node.IsLeaf() {
			return node
		}

		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		tmp := node
		node = t.min(tmp.right)
		node.right = t.deleteMin(tmp.right)
		node.left = tmp.left
		return tmp
	}

	return node
}
