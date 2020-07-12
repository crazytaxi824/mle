// Package set is not thread-safe
package set

import (
	"errors"
)

type uint16HashSet struct {
	elements map[uint16]struct{}
}

func NewUint16Set(n ...uint16) *uint16HashSet {
	var s uint16HashSet
	s.elements = make(map[uint16]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *uint16HashSet) Add(n uint16) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *uint16HashSet) Pop() (uint16, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New(ErrEmptySet)
}

// Delete element
func (s *uint16HashSet) Delete(element uint16) {
	delete(s.elements, element)
}

// Length of set
func (s *uint16HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *uint16HashSet) Range(fn func(element uint16) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *uint16HashSet) ToSlice() []uint16 {
	result := make([]uint16, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *uint16HashSet) Contains(n uint16) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *uint16HashSet) ContainsN(n []uint16) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *uint16HashSet) ContainsAny(n []uint16) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *uint16HashSet) Reset() {
	s.elements = make(map[uint16]struct{})
}

// Equal, elements
func (s *uint16HashSet) Equal(h *uint16HashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem uint16) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
