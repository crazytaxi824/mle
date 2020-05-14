package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceSet(t *testing.T) {
	var set = NewInterfaceSet()
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		set.Add(v)
	}
	ast := assert.New(t)
	ast.Equal(set.Contains(1, 2, 3), true)
}

func TestHashSetInt16(t *testing.T) {
	var s = []int16{1, 2, 3, 4, 5}
	var set = NewInt16Set(s...)

	ast := assert.New(t)
	ast.Equal(set.Contains(1, 2, 3), true)
	ast.Equal(set.Contains(1, 2, 6), false)

	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 9)), false)
	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 4, 5)), true)
}
