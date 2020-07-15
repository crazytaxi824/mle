// Package set is not thread-safe
package set

type int64HashSet struct {
	elements map[int64]struct{}
}

func NewInt64Set(n ...int64) *int64HashSet {
	var s int64HashSet
	s.elements = make(map[int64]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *int64HashSet) Add(n int64) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *int64HashSet) Pop() (int64, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, ErrEmptySet
}

// Delete element
func (s *int64HashSet) Delete(element int64) {
	delete(s.elements, element)
}

// Length of set
func (s *int64HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *int64HashSet) Range(fn func(element int64) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *int64HashSet) ToSlice() []int64 {
	result := make([]int64, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *int64HashSet) Contains(n int64) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *int64HashSet) ContainsN(n []int64) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *int64HashSet) ContainsAny(n []int64) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *int64HashSet) Reset() {
	s.elements = make(map[int64]struct{})
}

// Equal, elements
func (s *int64HashSet) Equal(h *int64HashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem int64) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
