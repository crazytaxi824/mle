package stack

import (
	"testing"
)

var elem = []int64{1, 2, 3, 4, 5, 6, 7}

func TestInt64Stack_Push(t *testing.T) {
	s := NewInt64ASCStack(false)
	for _, v := range elem {
		s.Push(v)
	}
	t.Log(s.Push(3))
	t.Log(s.Elements(), s.Len())
}

func TestInt64Stack_Pop(t *testing.T) {
	s := NewInt64ASCStack(false)
	for _, v := range elem {
		s.Push(v)
	}
	t.Log(s.Pop())
}

func TestInt64Stack_Search(t *testing.T) {
	s := NewInt64ASCStack(false)
	for _, v := range elem {
		s.Push(v)
	}
	if s.Search(7) != 1 {
		t.Fail()
	}
}
