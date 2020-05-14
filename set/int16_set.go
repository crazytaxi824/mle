// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetInt16 struct {
	elements map[int16]struct{}
}

func NewInt16Set(n ...int16) *hashSetInt16 {
	var s hashSetInt16
	s.elements = make(map[int16]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt16) Add(n int16) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt16) Pop() (int16, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetInt16) Delete(element int16) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetInt16) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetInt16) Range(fn func(element int16) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetInt16) ToSlice() []int16 {
	result := make([]int16, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetInt16) Contains(n int16, m ...int16) bool {
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
func (s *hashSetInt16) Reset() {
	s.elements = make(map[int16]struct{})
}

// Equal, elements
func (s *hashSetInt16) Equal(h *hashSetInt16) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem int16) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
