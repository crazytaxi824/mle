// Package set is not thread-safe
package set

import (
	"errors"
)

type uint32HashSet struct {
	elements map[uint32]struct{}
}

func NewUint32Set(n ...uint32) *uint32HashSet {
	var s uint32HashSet
	s.elements = make(map[uint32]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *uint32HashSet) Add(n uint32) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *uint32HashSet) Pop() (uint32, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New(ErrEmptySet)
}

// Delete element
func (s *uint32HashSet) Delete(element uint32) {
	delete(s.elements, element)
}

// Length of set
func (s *uint32HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *uint32HashSet) Range(fn func(element uint32) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *uint32HashSet) ToSlice() []uint32 {
	result := make([]uint32, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *uint32HashSet) Contains(n uint32) bool {
	_, ok := s.elements[n]
	return ok
}

func (s *uint32HashSet) ContainsN(n []uint32) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *uint32HashSet) ContainsAny(n []uint32) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *uint32HashSet) Reset() {
	s.elements = make(map[uint32]struct{})
}

// Equal, elements
func (s *uint32HashSet) Equal(h *uint32HashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem uint32) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
