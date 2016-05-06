package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a   = []string{"a"}
	c   = []string{"c"}
	cz  = []string{"c", "z"}
	ab  = []string{"a", "b"}
	abc = []string{"a", "b", "c"}
	abz = []string{"a", "b", "z"}
	xyz = []string{"x", "y", "z"}

	abcxyz = []string{"a", "b", "c", "x", "y", "z"}
)

func TestSet(t *testing.T) {
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
	cases := []struct {
		A        []string
		B        []string
		Expected []string
	}{
		{abc, xyz, append(xyz, abc...)},
		{abc, abz, append([]string{"a", "b", "c", "z"}, abc...)},
		{abc, nil, abc},
		{nil, abc, abc},
		{nil, nil, nil},
	}

	for _, tc := range cases {
		actual := NewSet(tc.A...).Union(NewSet(tc.B...))
		expected := NewSet(tc.Expected...)

		assert.Equal(t, expected, actual)
		if expected.Cardinality() > 0 {
			assert.True(t, !actual.IsEmpty())
		} else {
			assert.True(t, actual.IsEmpty())
		}
	}
}

func TestSetIntersection(t *testing.T) {
	cases := []struct {
		A        []string
		B        []string
		Expected []string
	}{
		{abc, xyz, nil},
		{abc, abz, ab},
		{abc, a, a},
		{a, abc, a},
		{nil, nil, nil},
	}

	for _, tc := range cases {
		actual := NewSet(tc.A...).Intersection(NewSet(tc.B...))
		expected := NewSet(tc.Expected...)

		assert.Equal(t, expected, actual)
		if expected.Cardinality() > 0 {
			assert.True(t, !actual.IsEmpty())
		} else {
			assert.True(t, actual.IsEmpty())
		}
	}
}

func TestSetDifference(t *testing.T) {
	cases := []struct {
		A        []string
		B        []string
		Expected []string
	}{
		{abc, xyz, abc},
		{abc, abz, c},
		{a, a, nil},
		{nil, nil, nil},
	}

	for _, tc := range cases {
		actual := NewSet(tc.A...).Difference(NewSet(tc.B...))
		expected := NewSet(tc.Expected...)

		assert.Equal(t, expected, actual)
		if expected.Cardinality() > 0 {
			assert.True(t, !actual.IsEmpty())
		} else {
			assert.True(t, actual.IsEmpty())
		}
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	cases := []struct {
		A        []string
		B        []string
		Expected []string
	}{
		{abc, xyz, abcxyz},
		{abc, abz, cz},
		{a, a, nil},
		{a, nil, a},
		{nil, a, a},
		{nil, nil, nil},
	}

	for _, tc := range cases {
		actual := NewSet(tc.A...).SymmetricDifference(NewSet(tc.B...))
		expected := NewSet(tc.Expected...)

		assert.Equal(t, expected, actual)
		if expected.Cardinality() > 0 {
			assert.True(t, !actual.IsEmpty())
		} else {
			assert.True(t, actual.IsEmpty())
		}
	}
}

func TestSetIsSubsetOf(t *testing.T) {
	cases := []struct {
		A        []string
		B        []string
		Expected bool
	}{
		{abc, abc, true},
		{ab, abc, true},
		{a, abc, true},
		{abc, xyz, false},
		{abc, ab, false},
		{nil, a, true},
		{nil, nil, true},
	}

	for _, tc := range cases {
		actual := NewSet(tc.A...).IsSubsetOf(NewSet(tc.B...))
		assert.Equal(t, tc.Expected, actual)
	}
}
