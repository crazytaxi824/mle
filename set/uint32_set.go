// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetUint32 struct {
	elements map[uint32]struct{}
}

func NewUint32Set(n ...uint32) *hashSetUint32 {
	var s hashSetUint32
	s.elements = make(map[uint32]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetUint32) Add(n uint32) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetUint32) Pop() (uint32, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetUint32) Delete(element uint32) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetUint32) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetUint32) Range(fn func(element uint32) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetUint32) ToSlice() []uint32 {
	result := make([]uint32, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetUint32) Contains(n uint32, m ...uint32) bool {
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
func (s *hashSetUint32) Reset() {
	s.elements = make(map[uint32]struct{})
}

// Equal, elements
func (s *hashSetUint32) Equal(h *hashSetUint32) bool {
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
