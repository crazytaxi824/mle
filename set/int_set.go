// Package set is not thread-safe
package set

type intHashSet struct {
	elements map[int]struct{}
}

func NewIntSet(n ...int) *intHashSet {
	var s intHashSet
	s.elements = make(map[int]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *intHashSet) Add(n int) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *intHashSet) Pop() (int, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, ErrEmptySet
}

// Delete element
func (s *intHashSet) Delete(element int) {
	delete(s.elements, element)
}

// Length of set
func (s *intHashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *intHashSet) Range(fn func(element int) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *intHashSet) ToSlice() []int {
	result := make([]int, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *intHashSet) Contains(n int) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *intHashSet) ContainsN(n []int) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *intHashSet) ContainsAny(n []int) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *intHashSet) Reset() {
	s.elements = make(map[int]struct{})
}

// Equal, elements
func (s *intHashSet) Equal(h *intHashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem int) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
