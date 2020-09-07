package set

import (
	"testing"

	"github.com/crazytaxi824/mle/util"
	"github.com/stretchr/testify/assert"
)

func TestHashSetInt16(t *testing.T) {
	var s = []int16{1, 2, 3, 4, 5, 1, 2, 3}
	var set = NewInt16Set(s...) // 去重

	ast := assert.New(t)
	ast.Equal(set.Contains(1), true)
	ast.Equal(set.Contains(7), false)

	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 9)), false)
	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 4, 5)), true)
}

func TestInterfaceHashSet(t *testing.T) {
	ast := assert.New(t)

	type person struct {
		name string
		age  int
	}

	fr := util.NewFakeRand()

	interSet := NewInterSet()

	for i := 0; i < 5; i++ {
		p := person{
			name: util.RandString(fr),
			age:  fr.Intn(60),
		}
		interSet.Add(p)
	}

	interSet.Add(person{
		name: "lq",
		age:  20,
	})
	interSet.Add(person{
		name: "zr",
		age:  20,
	})
	interSet.Add(person{
		name: "kk",
		age:  20,
	})

	// test add & len
	ast.Equal(interSet.Len(), 8)

	// test range
	interSet.Range(func(element interface{}) bool {
		t.Log("range: ", element)
		return true
	})

	// test delete
	interSet.Delete(person{
		name: "kk",
		age:  20,
	})
	ast.Equal(interSet.Len(), 7)

	// to slice
	t.Log(interSet.ToSlice())

	// contain
	ast.Equal(interSet.Contains(person{
		name: "lq",
		age:  20,
	}), true)

	ast.Equal(interSet.Contains(person{
		name: "kk",
		age:  21,
	}), false)

	// containsN
	ast.Equal(interSet.ContainsN([]interface{}{person{
		name: "lq",
		age:  20,
	}, person{
		name: "zr",
		age:  20,
	}}), true)

	ast.Equal(interSet.ContainsN([]interface{}{person{
		name: "lq",
		age:  20,
	}, person{
		name: "kk",
		age:  20,
	}}), false)

	// contains any
	ast.Equal(interSet.ContainsAny([]interface{}{person{
		name: "lq",
		age:  20,
	}, person{
		name: "kk",
		age:  20,
	}}), true)

	// equal
	tmp := interSet
	ast.Equal(interSet.Equal(tmp), true)

	// test pop
	t.Log(interSet.Pop())
	ast.Equal(interSet.Len(), 6)
}

// deep copy
func TestDeepCopy(t *testing.T) {
	a := NewInt16Set(1, 2, 3)
	b := a // shallow copy
	b.Add(4)
	if !a.Equal(b) {
		t.Fail()
	}

	c := a.Copy() // deep copy
	c.Add(5)
	if a.Equal(c) {
		t.Fail()
	}
}
