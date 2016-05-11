package bstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

 The default BST is perfectly balanced

        5
       / \
      3   7
     / \ / \
    2  4 6  8
*/

func defaultTree(t *testing.T) *Tree {
	ints := []int{5, 3, 4, 2, 7, 6, 8}
	tree := New()
	for _, key := range ints {
		tree.Insert(key)
	}
	return tree
}

func makeTree(ints []int) *Tree {
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
	cases := []struct {
		name     string
		expected []int
	}{
		{"in-order", []int{2, 3, 4, 5, 6, 7, 8}},
		{"pre-order", []int{5, 3, 2, 4, 7, 6, 8}},
		{"post-order", []int{2, 4, 3, 6, 8, 7, 5}},
	}

	for _, tc := range cases {
		tree := defaultTree(t)
		actual := make([]int, 0, tree.Len())
		collect := func(n *Node) { actual = append(actual, n.Key) }

		switch tc.name {
		case "in-order":
			tree.InOrder(collect)
		case "pre-order":
			tree.PreOrder(collect)
		case "post-order":
			tree.PostOrder(collect)
		}

		assert.Equal(t, tc.expected, actual)
	}
}

func TestTreeRemove(t *testing.T) {
	dtree := defaultTree(t)
	defaultMin := dtree.Min()
	defaultMax := dtree.Max()

	cases := []struct {
		name      string
		tree      *Tree
		removeKey int
		err       error
	}{
		{"empty", makeTree([]int{}), 1, ErrNodeNotFound},
		{"not-found-one-key", makeTree([]int{1}), 2, ErrNodeNotFound},
		{"root-one-key", makeTree([]int{1}), 1, nil},
		{"leaf-min-default", defaultTree(t), defaultMin.Key, nil},
		{"leaf-max-default", defaultTree(t), defaultMax.Key, nil},
		{"default-left", defaultTree(t), 3, nil},
		{"default-right", defaultTree(t), 7, nil},
		{"default-right", defaultTree(t), 7, nil},
	}

	for _, tc := range cases {
		origSize := tc.tree.Len()
		removed, err := tc.tree.Remove(tc.removeKey)

		assert.Equal(t, tc.err, err, tc.name)

		if tc.err == nil {
			assert.Equal(t, tc.removeKey, removed.Key, tc.name)
			assert.NoError(t, err, tc.name)
			// if we expected to remove a node then we should have -1 size
			assert.Equal(t, origSize-1, tc.tree.Len(), tc.name)
		} else {
			assert.Nil(t, removed, tc.name)
			assert.Error(t, err, tc.name)
			assert.Equal(t, origSize, tc.tree.Len(), tc.name)
		}
	}
}

func TestTreeHeight(t *testing.T) {
	tree := defaultTree(t)
	assert.Equal(t, 2, tree.Height())

	tree = New()
	tree.Insert(1)
	assert.Equal(t, 0, tree.Height())
}
