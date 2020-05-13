package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = []int16{1, 2, 3, 4, 5}

func TestSet(t *testing.T) {
	var set = NewInt16Set(s...)
	t.Log(set.Contains(1, 2, 3))
}

func TestHashSetInt16(t *testing.T) {
	var set = NewInt16Set(s...)

	ast := assert.New(t)
	ast.Equal(set.Contains(1, 2, 3), true)
	ast.Equal(set.Contains(1, 2, 6), false)

	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 9)), false)
	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 4, 5)), true)
}
