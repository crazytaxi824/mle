// Package set is not thread-safe
package set

type uint64HashSet struct {
	elements map[uint64]struct{}
}

func NewUint64Set(n ...uint64) *uint64HashSet {
	var s uint64HashSet
	s.elements = make(map[uint64]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *uint64HashSet) Add(n uint64) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *uint64HashSet) Pop() (uint64, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, ErrEmptySet
}

// Delete element
func (s *uint64HashSet) Delete(element uint64) {
	delete(s.elements, element)
}

// Length of set
func (s *uint64HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *uint64HashSet) Range(fn func(element uint64) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *uint64HashSet) ToSlice() []uint64 {
	result := make([]uint64, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *uint64HashSet) Contains(n uint64) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *uint64HashSet) ContainsN(n []uint64) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *uint64HashSet) ContainsAny(n []uint64) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *uint64HashSet) Reset() {
	s.elements = make(map[uint64]struct{})
}

// Equal, elements
func (s *uint64HashSet) Equal(h *uint64HashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem uint64) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}

// Copy, deep copy the elements
func (s *uint64HashSet) Copy() *uint64HashSet {
	var r = NewUint64Set()
	for k := range s.elements {
		r.Add(k)
	}
	return r
}
