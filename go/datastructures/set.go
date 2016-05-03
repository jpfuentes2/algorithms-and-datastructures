package datastructures

// Set implements an unordered set based on Go's map type
// The semantics for insert/lookup/remove are amortized O(1)
// which is a rough approximation (see https://medium.com/@ConnorPeet/go-maps-are-not-o-1-91c1e61110bf#.sp9v4ldi3)
type Set struct {
	data map[string]struct{}
}

// NewSet constructs a set
// Time O(n)
// Space O(n)
func NewSet(items ...string) *Set {
	s := &Set{data: make(map[string]struct{})}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Each iterates thru each element of the set
// Time O(n)
func (s *Set) Each(f func(string)) {
	for k := range s.data {
		f(k)
	}
}

// Add adds the element to the set
// Time O(1)
// Space O(1)
func (s *Set) Add(e string) {
	s.data[e] = struct{}{}
}

// Remove removes the element from the set if it exists
// Time O(1)
// Space O(1)
func (s *Set) Remove(e string) {
	delete(s.data, e)
}

// Contains returns true if the string is in set s
// Time O(1)
// Space O(1)
func (s *Set) Contains(e string) bool {
	_, ok := s.data[e]
	return ok
}

// IsEmpty is true if this is the empty set ∅
// Time O(1)
// Space O(1)
func (s *Set) IsEmpty() bool {
	return s.Cardinality() == 0
}

// Cardinality gives the cardinality/length of this set
// Time O(1)
// Space O(1)
func (s *Set) Cardinality() int {
	return s.Len()
}

// Len gives the cardinality/length of this set
// Time O(1)
// Space O(1)
func (s *Set) Len() int {
	return len(s.data)
}

// Union returns the union: s ∪ other
// Time O(n)
// Space O(n)
func (s *Set) Union(other *Set) *Set {
	out := NewSet()

	s.Each(func(str string) {
		out.Add(str)
	})

	other.Each(func(str string) {
		out.Add(str)
	})

	return out
}

// Intersection returns the intersection: s ∩ other
// Time O(n)
// Space O(n)
func (s *Set) Intersection(other *Set) *Set {
	out := NewSet()

	s.Each(func(str string) {
		if other.Contains(str) {
			out.Add(str)
		}
	})

	return out
}

// Difference is the set relative complement: other ∖ s
// Time O(n)
// Space O(n)
func (s *Set) Difference(other *Set) *Set {
	out := NewSet()

	s.Each(func(str string) {
		if !other.Contains(str) {
			out.Add(str)
		}
	})

	return out
}

// SymmetricDifference is the set of all elements in s and other which
// are not also in their intersection. More simply: any element in s not found in other and vice versa.
// Time O(n)
// Space O(n)
func (s *Set) SymmetricDifference(other *Set) *Set {
	return s
}

// IsSubsetOf tests whether s is a subset of other
// Time O(n)
// Space O(1)
func (s *Set) IsSubsetOf(other *Set) bool {
	for str := range s.data {
		if !other.Contains(str) {
			return false
		}
	}

	return true
}

func (s *Set) Product(other Set) *Set {
	return s
}
