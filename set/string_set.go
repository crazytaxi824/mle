// Package set is not thread-safe
package set

import (
	"errors"
)

type hashSetString struct {
	items map[string]struct{}
}

func NewStringSet(n ...string) *hashSetString {
	var s hashSetString
	s.items = make(map[string]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetString) Add(n string) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetString) Pop() (string, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return "", errors.New("empty set")
}

// Delete element
func (s *hashSetString) Delete(element string) {
	delete(s.items, element)
}

// Length of set
func (s *hashSetString) Len() int {
	return len(s.items)
}

// for range the set
func (s *hashSetString) Range(fn func(element string) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *hashSetString) ToSlice() []string {
	result := make([]string, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// Contains element
func (s *hashSetString) Contains(n string, m ...string) bool {
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
func (s *hashSetString) Reset() {
	s.items = make(map[string]struct{})
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
