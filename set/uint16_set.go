// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetUint16 struct {
	items map[uint16]struct{}
}

func NewUint16Set(n ...uint16) *hashSetUint16 {
	var s hashSetUint16
	s.items = make(map[uint16]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetUint16) Add(n uint16) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetUint16) Pop() (uint16, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetUint16) Delete(element uint16) {
	delete(s.items, element)
}

// Length of set
func (s *hashSetUint16) Len() int {
	return len(s.items)
}

// for range the set
func (s *hashSetUint16) Range(fn func(element uint16) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetUint16) ToSlice() []uint16 {
	result := make([]uint16, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetUint16) Contains(n uint16, m ...uint16) bool {
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
func (s *hashSetUint16) Reset() {
	s.items = make(map[uint16]struct{})
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
