// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetUint16 struct {
	elements map[uint16]struct{}
}

func NewUint16Set(n ...uint16) *hashSetUint16 {
	var s hashSetUint16
	s.elements = make(map[uint16]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetUint16) Add(n uint16) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetUint16) Pop() (uint16, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetUint16) Delete(element uint16) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetUint16) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetUint16) Range(fn func(element uint16) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetUint16) ToSlice() []uint16 {
	result := make([]uint16, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetUint16) Contains(n uint16, m ...uint16) bool {
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
func (s *hashSetUint16) Reset() {
	s.elements = make(map[uint16]struct{})
}

// Equal, elements
func (s *hashSetUint16) Equal(h *hashSetUint16) bool {
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
