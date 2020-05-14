// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetUint64 struct {
	items map[uint64]struct{}
}

func NewUint64Set(n ...uint64) *hashSetUint64 {
	var s hashSetUint64
	s.items = make(map[uint64]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetUint64) Add(n uint64) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetUint64) Pop() (uint64, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *hashSetUint64) Delete(element uint64) {
	delete(s.items, element)
}

// Length of set
func (s *hashSetUint64) Len() int {
	return len(s.items)
}

// for range the set
func (s *hashSetUint64) Range(fn func(element uint64) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetUint64) ToSlice() []uint64 {
	result := make([]uint64, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetUint64) Contains(n uint64, m ...uint64) bool {
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
func (s *hashSetUint64) Reset() {
	s.items = make(map[uint64]struct{})
}

// Equal, elements
func (s *hashSetUint64) Equal(h *hashSetUint64) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem uint64) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
