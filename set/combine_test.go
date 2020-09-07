package set

import (
	"testing"
)

func TestIntersect_Int16(t *testing.T) {
	a := NewInt16Set(1, 2, 3, 4, 5)
	b := NewInt16Set(4, 5, 6, 7, 8)

	r := Intersect().Int16(a, b)

	if !r.Equal(NewInt16Set(4, 5)) {
		t.Fail()
	}
}

func TestUnion_Int16(t *testing.T) {
	a := NewInt16Set(1, 2, 3, 4, 5)
	b := NewInt16Set(4, 5, 6, 7, 8)

	r := Union().Int16(a, b)

	if !r.Equal(NewInt16Set(1, 2, 3, 4, 5, 6, 7, 8)) {
		t.Fail()
	}
}

func TestSubtract_Int16(t *testing.T) {
	a := NewInt16Set(1, 2, 3, 4, 5)
	b := NewInt16Set(4, 5, 6, 7, 8)

	r := Subtract().Int16(a, b)

	if !r.Equal(NewInt16Set(1, 2, 3)) {
		t.Fail()
	}

	r = Subtract().Int16(b, a)

	if !r.Equal(NewInt16Set(6, 7, 8)) {
		t.Fail()
	}
}

func TestXor_Int16(t *testing.T) {
	a := NewInt16Set(1, 2, 3, 4, 5)
	b := NewInt16Set(4, 5, 6, 7, 8)

	r := XOR().Int16(a, b)

	if !r.Equal(NewInt16Set(1, 2, 3, 6, 7, 8)) {
		t.Fail()
	}
}
