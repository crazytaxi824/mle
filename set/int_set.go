// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetInt struct {
	items map[int]struct{}
}

func NewIntSet(n ...int) *hashSetInt {
	var s hashSetInt
	s.items = make(map[int]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt) Add(n int) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt) Pop() (int, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetInt) Delete(element int) {
	delete(s.items, element)
}

// Length of set
func (s *hashSetInt) Len() int {
	return len(s.items)
}

// for range the set
func (s *hashSetInt) Range(fn func(element int) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetInt) ToSlice() []int {
	result := make([]int, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetInt) Contains(n int, m ...int) bool {
	if _, ok := s.items[n]; !ok {
		return false
	}

	for _, v := range m {
		if _, ok := s.items[v]; !ok {
			return false
		}
	}
	return true
}

// Reset the set
func (s *hashSetInt) Reset() {
	s.items = make(map[int]struct{})
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
