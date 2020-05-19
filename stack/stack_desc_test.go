package stack

import (
	"testing"
)

var elemDESC = []int64{7, 6, 5, 4, 3, 2, 1}

func TestInt64DESCStack_Push(t *testing.T) {
	s := NewInt64DESCStack(false)
	for _, v := range elemDESC {
		s.Push(v)
	}
	t.Log(s.Push(3))
	t.Log(s.Elements(), s.Len())
}

func TestInt64DESCStack_Pop(t *testing.T) {
	s := NewInt64DESCStack(false)
	for _, v := range elemDESC {
		s.Push(v)
	}
	t.Log(s.Pop())
}

func TestInt64DESCStack_Search(t *testing.T) {
	s := NewInt64DESCStack(false)
	for _, v := range elemDESC {
		s.Push(v)
	}
	if s.Search(7) != 7 {
		t.Fail()
	}
}
