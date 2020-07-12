// Package set is not thread-safe
package set

import (
	"errors"
)

type runeHashSet struct {
	elements map[rune]struct{}
}

func NewRuneSet(n ...rune) *runeHashSet {
	var s runeHashSet
	s.elements = make(map[rune]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *runeHashSet) Add(n rune) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *runeHashSet) Pop() (rune, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return 0, errors.New(ErrEmptySet)
}

// Delete element
func (s *runeHashSet) Delete(element rune) {
	delete(s.elements, element)
}

// Length of set
func (s *runeHashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *runeHashSet) Range(fn func(element rune) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *runeHashSet) ToSlice() []rune {
	result := make([]rune, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *runeHashSet) Contains(n rune) bool {
	_, ok := s.elements[n]
	return ok
}

// Contains all elements
func (s *runeHashSet) ContainsN(n []rune) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *runeHashSet) ContainsAny(n []rune) bool {
	for _, v := range n {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *runeHashSet) Reset() {
	s.elements = make(map[rune]struct{})
}

// Equal, elements
func (s *runeHashSet) Equal(h *runeHashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem rune) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
