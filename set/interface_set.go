// Package set is not thread-safe
package set

import (
	"errors"
)

type interfaceHashSet struct {
	elements map[interface{}]struct{}
}

func NewInterSet(n ...interface{}) *interfaceHashSet {
	var s interfaceHashSet
	s.elements = make(map[interface{}]struct{})
	for _, v := range n {
		s.elements[v] = struct{}{}
	}
	return &s
}

// add element
func (s *interfaceHashSet) Add(n interface{}) {
	s.elements[n] = struct{}{}
}

// Pop random element
func (s *interfaceHashSet) Pop() (interface{}, error) {
	for k := range s.elements {
		delete(s.elements, k)
		return k, nil
	}
	return nil, errors.New(ErrEmptySet)
}

// Delete element
func (s *interfaceHashSet) Delete(element interface{}) {
	delete(s.elements, element)
}

// Length of set
func (s *interfaceHashSet) Len() int {
	return len(s.elements)
}

// for range the set
func (s *interfaceHashSet) Range(fn func(element interface{}) bool) {
	for k := range s.elements {
		if !fn(k) {
			return
		}
	}
}

// ToSlice return slice
func (s *interfaceHashSet) ToSlice() []interface{} {
	result := make([]interface{}, len(s.elements))
	var count int
	for k := range s.elements {
		result[count] = k
		count++
	}
	return result
}

// Contains element
func (s *interfaceHashSet) Contains(element interface{}) bool {
	_, ok := s.elements[element]
	return ok
}

// Contains all of the elements
func (s *interfaceHashSet) ContainsN(elements []interface{}) bool {
	for _, v := range elements {
		if _, ok := s.elements[v]; !ok {
			return false
		}
	}
	return true
}

func (s *interfaceHashSet) ContainsAny(elements []interface{}) bool {
	for _, v := range elements {
		if _, ok := s.elements[v]; ok {
			return true
		}
	}
	return false
}

// Reset the set
func (s *interfaceHashSet) Reset() {
	s.elements = make(map[interface{}]struct{})
}

// Equal, elements
func (s *interfaceHashSet) Equal(h *interfaceHashSet) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(elem interface{}) bool {
		if s.Contains(elem) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
