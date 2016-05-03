package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Parallel()

	letters := []string{"a", "b", "c"}

	assert.True(t, NewSet().IsEmpty())

	set := NewSet(letters...)
	assert.False(t, set.IsEmpty())

	for _, l := range letters {
		assert.True(t, set.Contains(l))
	}

	assert.False(t, set.Contains("d"))
	assert.False(t, set.IsEmpty())

	for _, l := range letters {
		set.Remove(l)
		assert.False(t, set.Contains(l))
	}

	assert.True(t, set.IsEmpty())
}

func TestSetUnion(t *testing.T) {
	set1 := NewSet("a", "b", "c")
	set2 := NewSet("x", "y", "z")
	union := NewSet("a", "b", "c", "x", "y", "z")

	assert.Equal(t, union, set1.Union(set2))

	set1 = NewSet("a", "b", "c")
	set2 = NewSet("a", "b", "z")
	union = NewSet("a", "b", "c", "z")

	assert.Equal(t, union, set1.Union(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Union(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetIntersection(t *testing.T) {
	set1 := NewSet("a", "b", "c")
	set2 := NewSet("x", "y", "z")
	intersect := NewSet()

	assert.Equal(t, intersect, set1.Intersection(set2))

	set1 = NewSet("a", "b", "c")
	set2 = NewSet("a", "b", "z")
	intersect = NewSet("a", "b")

	assert.Equal(t, intersect, set1.Intersection(set2))

	// handle case where one set has less elements
	set1 = NewSet("a")
	set2 = NewSet("a", "b", "z")
	intersect = NewSet("a")

	assert.Equal(t, intersect, set1.Intersection(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Intersection(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetDifference(t *testing.T) {
	set1 := NewSet("a", "b", "c")
	set2 := NewSet("x", "y", "z")
	diff := NewSet("a", "b", "c")

	assert.Equal(t, diff, set1.Difference(set2))

	set1 = NewSet("a", "b", "c")
	set2 = NewSet("a", "b", "z")
	diff = NewSet("c")

	assert.Equal(t, diff, set1.Difference(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Difference(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetSymmetricDifference(t *testing.T) {
	set1 := NewSet("a", "b", "c")
	set2 := NewSet("x", "y", "z")
	symm := NewSet("a", "b", "c", "x", "y", "z")

	assert.Equal(t, symm, set1.SymmetricDifference(set2))

	set1 = NewSet("a", "b", "1")
	set2 = NewSet("a", "b", "2")
	symm = NewSet("1", "2")

	assert.Equal(t, symm, set1.SymmetricDifference(set2))

	set1 = NewSet("a", "b")
	set2 = NewSet("a", "b")
	symm = NewSet()

	assert.Equal(t, symm, set1.SymmetricDifference(set2))
	assert.True(t, set1.SymmetricDifference(set2).IsEmpty())
}

func TestSetIsSubsetOf(t *testing.T) {
	set1 := NewSet("a", "b", "c")
	set2 := NewSet("x", "y", "z")

	assert.False(t, set1.IsSubsetOf(set2))
	assert.False(t, set2.IsSubsetOf(set1))
	// set is subset of itself
	assert.True(t, set1.IsSubsetOf(set1))
	assert.True(t, set2.IsSubsetOf(set2))

	set1 = NewSet("a", "b", "c")
	set2 = NewSet("a", "b", "c", "d")

	assert.True(t, set1.IsSubsetOf(set2))
	assert.False(t, set2.IsSubsetOf(set1))

	set1 = NewSet("a", "b", "c")
	set2 = NewSet("a", "b", "d")

	assert.False(t, set1.IsSubsetOf(set2))
	assert.False(t, set2.IsSubsetOf(set1))
}
