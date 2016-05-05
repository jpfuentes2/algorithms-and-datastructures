package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBSTree(t *testing.T) {
	tree := NewBSTree()
	assert.True(t, tree.IsEmpty())
	assert.Equal(t, 0, tree.Len())
}
