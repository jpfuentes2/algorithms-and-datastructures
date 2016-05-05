package datastructures

// BSTreeNode is a binary search tree node
type BSTreeNode struct {
	key      int
	value    int
	numNodes int

	left  *BSTreeNode
	right *BSTreeNode
}

// BSTree is a binary search tree
type BSTree struct {
	root *BSTreeNode
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func (t *BSTree) IsEmpty() bool {
	return t.Len() == 0
}

func (t *BSTree) Len() int {
	if t.root == nil {
		return 0
	}

	return t.root.numNodes
}

func (t *BSTree) Insert() {

}

func (t *BSTree) Remove() {

}
