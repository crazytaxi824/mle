// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetString struct {
	elements map[string]struct{}
}

func NewStringSet(n ...string) *hashSetString {
	var s hashSetString
	s.elements = make(map[string]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetString) Add(n string) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *hashSetString) Pop() (string, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return "", errors.New("empty set")
}

// Delete element
func (s *hashSetString) Delete(element string) {
	delete(s.elements, element)
}

// Length of set
func (s *hashSetString) Len() int {
	return len(s.elements)
}

// for range the set
func (s *hashSetString) Range(fn func(element string) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetString) ToSlice() []string {
	result := make([]string, len(s.elements))
	var count int
	for i := range s.elements {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetString) Contains(n string, m ...string) bool {
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
func (s *hashSetString) Reset() {
	s.elements = make(map[string]struct{})
}

// Equal, elements
func (s *hashSetString) Equal(h *hashSetString) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem string) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
