// Package set 不是线程安全的
package set

type hashSet struct {
	set map[int16]struct{}
}

func New() hashSet {
	var s hashSet
	s.set = make(map[int16]struct{})
	return s
}

func NewFrom(n []int16) hashSet {
	var s hashSet
	s.set = make(map[int16]struct{})
	for k := range n {
		s.set[n[k]] = struct{}{}
	}
	return s
}

// add element
func (s *hashSet) Add(n int16) {
	s.set[n] = struct{}{}
}

// Pop random element
func (s *hashSet) Pop() int16 {
	for k := range s.set {
		delete(s.set, k)
		return k
	}
	panic("empty set")
}

// len 长度
func (s *hashSet) Len() int {
	return len(s.set)
}

// 循环，返回 false 终止循环
func (s *hashSet) Range(fn func(key int16) bool) {
	for k := range s.set {
		if !fn(k) {
			return
		}
	}
}
