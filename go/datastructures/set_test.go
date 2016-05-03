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
	set1 := NewSet([]string{"a", "b", "c"}...)
	set2 := NewSet([]string{"x", "y", "z"}...)
	union := NewSet([]string{"a", "b", "c", "x", "y", "z"}...)

	assert.Equal(t, union, set1.Union(set2))

	set1 = NewSet([]string{"a", "b", "c"}...)
	set2 = NewSet([]string{"a", "b", "z"}...)
	union = NewSet([]string{"a", "b", "c", "z"}...)

	assert.Equal(t, union, set1.Union(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Union(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetIntersection(t *testing.T) {
	set1 := NewSet([]string{"a", "b", "c"}...)
	set2 := NewSet([]string{"x", "y", "z"}...)
	intersect := NewSet()

	assert.Equal(t, intersect, set1.Intersection(set2))

	set1 = NewSet([]string{"a", "b", "c"}...)
	set2 = NewSet([]string{"a", "b", "z"}...)
	intersect = NewSet([]string{"a", "b", "c", "z"}...)

	assert.Equal(t, intersect, set1.Union(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Intersection(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetDifference(t *testing.T) {
	set1 := NewSet([]string{"a", "b", "c"}...)
	set2 := NewSet([]string{"x", "y", "z"}...)
	diff := NewSet([]string{"a", "b", "c"}...)

	assert.Equal(t, diff, set1.Difference(set2))

	set1 = NewSet([]string{"a", "b", "c"}...)
	set2 = NewSet([]string{"a", "b", "z"}...)
	diff = NewSet([]string{"c"}...)

	assert.Equal(t, diff, set1.Difference(set2))

	empty := NewSet()
	assert.Equal(t, empty, empty.Difference(empty))
	assert.True(t, empty.IsEmpty())
}

func TestSetSymmetricDifference(t *testing.T) {
	assert.Fail(t, ":(")
}

func TestSetProduct(t *testing.T) {
	assert.Fail(t, ":(")
}

func TestSetIsSubsetOf(t *testing.T) {
	assert.Fail(t, ":(")
}
