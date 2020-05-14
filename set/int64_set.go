// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetInt64 struct {
	elements map[int64]struct{}
}

func NewInt64Set(n ...int64) *hashSetInt64 {
	var s hashSetInt64
	s.elements = make(map[int64]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt64) Add(n int64) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt64) Pop() (int64, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetInt64) Delete(element int64) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetInt64) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetInt64) Range(fn func(element int64) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetInt64) ToSlice() []int64 {
	result := make([]int64, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetInt64) Contains(n int64, m ...int64) bool {
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
func (s *hashSetInt64) Reset() {
	s.elements = make(map[int64]struct{})
}

// Equal, elements
func (s *hashSetInt64) Equal(h *hashSetInt64) bool {
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
