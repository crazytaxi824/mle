// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetInt struct {
	elements map[int]struct{}
}

func NewIntSet(n ...int) *hashSetInt {
	var s hashSetInt
	s.elements = make(map[int]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt) Add(n int) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt) Pop() (int, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetInt) Delete(element int) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetInt) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetInt) Range(fn func(element int) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetInt) ToSlice() []int {
	result := make([]int, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetInt) Contains(n int, m ...int) bool {
	if _, ok := s.elements[n]; !ok {
		return false
	}

	for _, v := range m {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

// Reset the set
func (s *hashSetInt) Reset() {
	s.elements = make(map[int]struct{})
}

// Equal, elements
func (s *hashSetInt) Equal(h *hashSetInt) bool {
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
