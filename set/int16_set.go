// Package set is not thread-safe
package set

import (
	"errors"
)

type int16HashSet struct {
	elements map[int16]struct{}
}

func NewInt16Set(n ...int16) *int16HashSet {
	var s int16HashSet
	s.elements = make(map[int16]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *int16HashSet) Add(n int16) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *int16HashSet) Pop() (int16, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New("empty set")
}

// Delete element
func (s *int16HashSet) Delete(element int16) {
	delete(s.elements, element)
}

// Length of set
func (s *int16HashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *int16HashSet) Range(fn func(element int16) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *int16HashSet) ToSlice() []int16 {
	result := make([]int16, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *int16HashSet) Contains(n int16) bool {
	_, ok := s.elements[n]
	return ok
}

func (s *int16HashSet) ContainsN(n []int16) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *int16HashSet) ContainsAny(n []int16) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *int16HashSet) Reset() {
	s.elements = make(map[int16]struct{})
}

// Equal, elements
func (s *int16HashSet) Equal(h *int16HashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem int16) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
