package bstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultTree(t *testing.T) *Tree {
	ints := []int{5, 4, 6, 3, 7, 2, 8}
	tree := New()
	for _, key := range ints {
		tree.Insert(key)
	}
	return tree
}

func TestEmptyTree(t *testing.T) {
	tree := New()
	assert.True(t, tree.IsEmpty())
	assert.Equal(t, 0, tree.Len())
	assert.Nil(t, tree.Min())
	assert.Nil(t, tree.Max())
}

func TestTree(t *testing.T) {
	tree := defaultTree(t)

	assert.Equal(t, 7, tree.Len())
	assert.True(t, tree.Contains(8))
	assert.True(t, tree.Contains(2))
	assert.False(t, tree.Contains(1))
	assert.False(t, tree.Contains(9))

}

func TestTreeMinMax(t *testing.T) {
	cases := []struct {
		ints []int
		min  int
		max  int
	}{
		{[]int{5}, 5, 5},
		{[]int{5, 6}, 5, 6},
		{[]int{6, 5}, 5, 6},
		{[]int{}, 0, 0},
	}

	for _, tc := range cases {
		tree := New()
		for _, i := range tc.ints {
			tree.Insert(i)
		}

		if len(tc.ints) > 0 {
			assert.Equal(t, tc.min, tree.Min().Key)
			assert.Equal(t, tc.max, tree.Max().Key)
		} else {
			assert.Nil(t, tree.Min())
			assert.Nil(t, tree.Max())
		}
	}
}

func TestTreeDuplicateInsert(t *testing.T) {
	tree := New()

	err := tree.Insert(5)
	assert.NoError(t, err)

	err = tree.Insert(5)
	assert.Equal(t, ErrDuplicateKey, err)
	assert.Equal(t, 1, tree.Len())
}

func TestTreeOrdering(t *testing.T) {
	order := []int{}
	tree := defaultTree(t)

	tree.InOrder(func(node *Node) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{2, 3, 4, 5, 6, 7, 8})

	order = order[:0]
	tree.PreOrder(func(node *Node) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{5, 4, 3, 2, 6, 7, 8})

	order = order[:0]
	tree.PostOrder(func(node *Node) {
		order = append(order, node.Key)
	})
	assert.Equal(t, order, []int{2, 3, 4, 8, 7, 6, 5})
}

func TestTreeRemove(t *testing.T) {
	// removing from an empty tree or nil node is a noop
	tree := New()
	assert.Error(t, tree.Remove(&Node{}))
	tree.Insert(1)
	assert.Error(t, tree.Remove(nil))

	tree = defaultTree(t)
	// remove leaf nodes
	assert.NoError(t, tree.Remove(tree.Min()))
	assert.NoError(t, tree.Remove(tree.Max()))

	// tree = defaultTree(t)
	// // remove leaf nodes
	// assert.NoError(t, tree.Remove(tree.Min()))
	// assert.NoError(t, tree.Remove(tree.Max()))

	// assert.Fail(t, ":(")
}

func TestTreeHeight(t *testing.T) {
	tree := defaultTree(t)
	assert.Equal(t, 3, tree.Height())

	tree = New()
	tree.Insert(1)
	assert.Equal(t, 0, tree.Height())
}
