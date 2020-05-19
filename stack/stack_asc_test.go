package stack

import (
	"reflect"
	"testing"
)

var elemASC = []int64{1, 2, 3, 4, 5, 6, 7}

func initASCStack(allowDupl bool) *int64ASCStack {
	s := NewInt64ASCStack(allowDupl)
	for _, v := range elemASC {
		s.Push(v)
	}
	return s
}

func TestInt64ASCStack_Push(t *testing.T) {
	s := initASCStack(false)

	n, res := s.Push(3)

	if n != 5 {
		t.Fail()
	}

	if reflect.DeepEqual(res, []int{3, 4, 5, 6, 7}) {
		t.Fail()
	}

	if reflect.DeepEqual(s.Elements(), []int{1, 2, 3}) {
		t.Fail()
	}

	// allow duplicated elements
	s2 := initASCStack(true)
	n, res = s2.Push(3)

	if n != 4 {
		t.Fail()
	}

	if reflect.DeepEqual(res, []int{4, 5, 6, 7}) {
		t.Fail()
	}

	if reflect.DeepEqual(s.Elements(), []int{1, 2, 3, 3}) {
		t.Fail()
	}
}

func TestInt64ASCStack_Pop(t *testing.T) {
	s := initASCStack(false)

	i, err := s.Pop()
	if err != nil {
		t.Error(err)
		return
	}

	if i != 7 {
		t.Fail()
	}

	if s.Len() != 6 {
		t.Fail()
	}
}

func TestInt64ASCStack_Peek(t *testing.T) {
	s := initASCStack(false)

	i, err := s.Peek()
	if err != nil {
		t.Error(err)
		return
	}
	if i != 7 {
		t.Fail()
	}

	if s.Len() != 7 {
		t.Fail()
	}
}

func TestInt64ASCStack_Search(t *testing.T) {
	s := initASCStack(false)

	if s.Search(7) != 1 {
		t.Fail()
	}
}

func TestInt64ASCStack_IsEmpty(t *testing.T) {
	s := initASCStack(false)
	if s.IsEmpty() {
		t.Fail()
	}
}

func TestInt64ASCStack_Reset(t *testing.T) {
	s := initASCStack(false)
	if s.IsEmpty() {
		t.Fail()
	}

	s.Reset()
	if !s.IsEmpty() {
		t.Fail()
	}
}
