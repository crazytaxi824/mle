package stack

import (
	"testing"
)

// new Stack
var elem = []int64{1, 2, 3, 4, 5, 4, 3, 2, 3}

func TestNormalInt64Stack(t *testing.T) {
	normalStack := NewInt64Stack(nil)
	for _, v := range elem {
		normalStack.Push(v)
	}

	t.Log(normalStack.Elements())

	normalStack2 := NewInt64Stack(&Option{
		AllowDuplicate: false,
		StackType:      NormalStack,
	})
	for _, v := range elem {
		normalStack2.Push(v)
	}

	t.Log(normalStack2.Elements())
}

func TestASCInt64Stack(t *testing.T) {
	ascStack := NewInt64Stack(&Option{
		AllowDuplicate: true,
		StackType:      ASCStack,
	})

	for _, v := range elem {
		ascStack.Push(v)
	}

	t.Log(ascStack.Elements())

	ascStack2 := NewInt64Stack(&Option{
		AllowDuplicate: false,
		StackType:      ASCStack,
	})

	for _, v := range elem {
		ascStack2.Push(v)
	}

	t.Log(ascStack2.Elements())
}

func TestDESCInt64Stack(t *testing.T) {
	descStack := NewInt64Stack(&Option{
		AllowDuplicate: true,
		StackType:      DESCStack,
	})
	for _, v := range elem {
		descStack.Push(v)
	}

	t.Log(descStack.Elements())

	descStack2 := NewInt64Stack(&Option{
		AllowDuplicate: false,
		StackType:      DESCStack,
	})
	for _, v := range elem {
		descStack2.Push(v)
	}

	t.Log(descStack2.Elements())
}

// ASC stack
var elemASC = []int64{1, 2, 3, 4, 5, 6, 7}

func initASCStack(allowDupl bool) *int64Stack {
	asc := NewInt64Stack(&Option{
		AllowDuplicate: allowDupl,
		StackType:      ASCStack,
	})

	for _, v := range elemASC {
		asc.Push(v)
	}

	return asc
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

	if s.Search(2) != s.searchASCDichotomy(2) {
		t.Fail()
	}
}

func TestInt64ASCStack_IsEmpty(t *testing.T) {
	s := initASCStack(true)
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

func TestInt64ASCStack_Range(t *testing.T) {
	s := initASCStack(false)

	var count int64
	s.Range(func(element int64) bool {
		count++
		if element != count {
			t.Fail()
			return false
		}
		return true
	})
}

// DESC stack

var elemDESC = []int64{7, 6, 5, 4, 3, 2, 1}

func initDESCStack(allowDupl bool) *int64Stack {
	asc := NewInt64Stack(&Option{
		AllowDuplicate: allowDupl,
		StackType:      DESCStack,
	})

	for _, v := range elemDESC {
		asc.Push(v)
	}

	return asc
}

func TestInt64DESCStack_Pop(t *testing.T) {
	s := initDESCStack(false)

	i, err := s.Pop()
	if err != nil {
		t.Error(err)
		return
	}

	if i != 1 {
		t.Fail()
	}

	if s.Len() != 6 {
		t.Fail()
	}
}

func TestInt64DESCStack_Peek(t *testing.T) {
	s := initDESCStack(false)

	i, err := s.Peek()
	if err != nil {
		t.Error(err)
		return
	}
	if i != 1 {
		t.Fail()
	}

	if s.Len() != 7 {
		t.Fail()
	}
}

func TestInt64DESCStack_Search(t *testing.T) {
	s := initDESCStack(false)

	if s.Search(1) != 1 {
		t.Fail()
	}

	if s.Search(4) != s.searchDESCDichotomy(4) {
		t.Fail()
	}
}

func TestInt64Stack_Search2(t *testing.T) {
	e := []int64{5, 4, 3, 3, 3}

	s := NewInt64Stack(&Option{
		AllowDuplicate: true,
		StackType:      DESCStack,
	})

	for _, v := range e {
		s.Push(v)
	}

	if s.Search(1) != -1 {
		t.Fail()
	}

	if s.Search(3) != 1 {
		t.Fail()
	}

	if s.Search(4) != 4 {
		t.Fail()
	}
}

func TestInt64DESCStack_IsEmpty(t *testing.T) {
	s := initDESCStack(true)
	if s.IsEmpty() {
		t.Fail()
	}
}

func TestInt64DESCStack_Reset(t *testing.T) {
	s := initDESCStack(false)
	if s.IsEmpty() {
		t.Fail()
	}

	s.Reset()
	if !s.IsEmpty() {
		t.Fail()
	}
}

func TestInt64DESCStack_Range(t *testing.T) {
	s := initDESCStack(false)

	var count = int64(7)
	s.Range(func(element int64) bool {
		if element != count {
			t.Fail()
			return false
		}
		count--
		return true
	})
}

// bench
func BenchmarkASCSearch(b *testing.B) {
	ascStack := NewInt64Stack(&Option{
		AllowDuplicate: false,
		StackType:      ASCStack,
	})

	for i := 0; i < 1000; i++ {
		ascStack.Push(int64(i))
	}

	for i := 0; i < b.N; i++ {
		ascStack.searchASCDichotomy(501)
	}
	b.ReportAllocs()
}

func BenchmarkNormalSearch(b *testing.B) {
	ascStack := NewInt64Stack(&Option{
		AllowDuplicate: false,
		StackType:      ASCStack,
	})

	for i := 0; i < 1000; i++ {
		ascStack.Push(int64(i))
	}

	for i := 0; i < b.N; i++ {
		ascStack.searchNormalStack(501)
	}
	b.ReportAllocs()
}
