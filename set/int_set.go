// Package set 不是线程安全的
package set

import (
	"errors"
)

type hashSetInt16 struct {
	items map[int16]struct{}
}

func NewInt16Set(n ...int16) *hashSetInt16 {
	var s hashSetInt16
	s.items = make(map[int16]struct{})
	for _, v := range n {
		s.items[v] = struct{}{}
	}
	return &s
}

// add element
func (s *hashSetInt16) Add(n int16) {
	s.items[n] = struct{}{}
}

// Pop random element
func (s *hashSetInt16) Pop() (int16, error) {
	for k := range s.items {
		delete(s.items, k)
		return k, nil
	}
	return 0, errors.New("empty items")
}

// Delete element
func (s *hashSetInt16) Delete(n int16) {
	delete(s.items, n)
}

// Len 长度
func (s *hashSetInt16) Len() int {
	return len(s.items)
}

// 循环，返回 false 终止循环
func (s *hashSetInt16) Range(fn func(key int16) bool) {
	for k := range s.items {
		if !fn(k) {
			return
		}
	}
}

// ToSlice 转成 []int16
func (s *hashSetInt16) ToSlice() []int16 {
	result := make([]int16, len(s.items))
	var count int
	for i := range s.items {
		result[count] = i
		count++
	}
	return result
}

// ContainsLong element
func (s *hashSetInt16) Contains(n int16, m ...int16) bool {
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

// IsEmpty, true - empty items
func (s *hashSetInt16) IsEmpty() bool {
	return len(s.items) == 0
}

// Reset the items
func (s *hashSetInt16) Reset() {
	s.items = make(map[int16]struct{})
}

// Equal, elements
func (s *hashSetInt16) Equal(h *hashSetInt16) bool {
	if s.Len() != h.Len() {
		return false
	}

	var mark = true
	h.Range(func(key int16) bool {
		if s.Contains(key) {
			return true
		}
		mark = false
		return false
	})
	return mark
}
