// Package set is not thread-safe
package set

import (
	"errors"
)

type int32HashSet struct {
	elements map[int32]struct{}
}

func NewInt32Set(n ...int32) *int32HashSet {
	var s int32HashSet
	s.elements = make(map[int32]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *int32HashSet) Add(n int32) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *int32HashSet) Pop() (int32, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New(ErrEmptySet)
}

// Delete element
func (s *int32HashSet) Delete(element int32) {
	delete(s.elements, element)
}

// Length of set
func (s *int32HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *int32HashSet) Range(fn func(element int32) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *int32HashSet) ToSlice() []int32 {
	result := make([]int32, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *int32HashSet) Contains(n int32) bool {
	_, ok := s.elements[n]
	return ok
}

func (s *int32HashSet) ContainsN(n []int32) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *int32HashSet) ContainsAny(n []int32) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *int32HashSet) Reset() {
	s.elements = make(map[int32]struct{})
}

// Equal, elements
func (s *int32HashSet) Equal(h *int32HashSet) bool {
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
