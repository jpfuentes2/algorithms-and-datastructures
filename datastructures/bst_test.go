package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultTree(t *testing.T) *BSTree {
	ints := []int{5, 4, 6, 3, 7, 2, 8}
	tree := NewBSTree()
	for _, key := range ints {
		tree.Insert(key)
	}
	return tree
}

func TestEmptyBSTree(t *testing.T) {
	tree := NewBSTree()
	assert.True(t, tree.IsEmpty())
	assert.Equal(t, 0, tree.Len())
	assert.Nil(t, tree.Min())
	assert.Nil(t, tree.Max())
}

func TestBSTree(t *testing.T) {
	tree := defaultTree(t)

	assert.Equal(t, 7, tree.Len())
	assert.True(t, tree.Contains(8))
	assert.True(t, tree.Contains(2))
	assert.False(t, tree.Contains(1))
	assert.False(t, tree.Contains(9))

	assert.Equal(t, 2, tree.Min().Key)
	assert.Equal(t, 8, tree.Max().Key)
}

func TestBSTreeDuplicateInsert(t *testing.T) {
	tree := NewBSTree()

	err := tree.Insert(5)
	assert.NoError(t, err)

	err = tree.Insert(5)
	assert.Equal(t, ErrDuplicateKey, err)
	assert.Equal(t, 1, tree.Len())
}

func TestBSTreeOrdering(t *testing.T) {
	order := []int{}
	tree := defaultTree(t)

	tree.InOrder(func(node *BSTreeNode) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{2, 3, 4, 5, 6, 7, 8})

	order = order[:0]
	tree.PreOrder(func(node *BSTreeNode) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{5, 4, 3, 2, 6, 7, 8})

	order = order[:0]
	tree.PostOrder(func(node *BSTreeNode) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{2, 3, 4, 8, 7, 6, 5})
}

func TestBSTreeRemove(t *testing.T) {
	_ = defaultTree(t)
	assert.Fail(t, ":(")
}

func TestBSTreeHeight(t *testing.T) {
	tree := defaultTree(t)
	assert.Equal(t, 3, tree.Height())

	tree = NewBSTree()
	tree.Insert(1)
	assert.Equal(t, 0, tree.Height())
}
