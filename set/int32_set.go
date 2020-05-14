// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetInt32 struct {
	items map[int32]struct{}
}

func NewInt32Set(n ...int32) *hashSetInt32 {
	var s hashSetInt32
	s.items = make(map[int32]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt32) Add(n int32) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt32) Pop() (int32, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetInt32) Delete(element int32) {
	delete(s.items, element)
}

// Length of set
func (s *hashSetInt32) Len() int {
	return len(s.items)
}

// for range the set
func (s *hashSetInt32) Range(fn func(element int32) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetInt32) ToSlice() []int32 {
	result := make([]int32, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetInt32) Contains(n int32, m ...int32) bool {
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
func (s *hashSetInt32) Reset() {
	s.items = make(map[int32]struct{})
}

// Equal, elements
func (s *hashSetInt32) Equal(h *hashSetInt32) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem int32) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
