// Package set is not thread-safe
package set

type stringHashSet struct {
	elements map[string]struct{}
}

func NewStringSet(n ...string) *stringHashSet {
	var s stringHashSet
	s.elements = make(map[string]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *stringHashSet) Add(n string) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *stringHashSet) Pop() (string, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return "", ErrEmptySet
}

// Delete element
func (s *stringHashSet) Delete(element string) {
	delete(s.elements, element)
}

// Length of set
func (s *stringHashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *stringHashSet) Range(fn func(element string) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *stringHashSet) ToSlice() []string {
	result := make([]string, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *stringHashSet) Contains(n string) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *stringHashSet) ContainsN(n []string) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *stringHashSet) ContainsAny(n []string) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *stringHashSet) Reset() {
	s.elements = make(map[string]struct{})
}

// Equal, elements
func (s *stringHashSet) Equal(h *stringHashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem string) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}

// Copy, deep copy the elements
func (s *stringHashSet) Copy() *stringHashSet {
	var r = NewStringSet()
	for k := range s.elements {
		r.Add(k)
	}
	return r
}
