package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSetInt16(t *testing.T) {
	var s = []int16{1, 2, 3, 4, 5}
	var set = NewInt16Set(s...)

	ast := assert.New(t)
	ast.Equal(set.Contains(1, 2, 3), true)
	ast.Equal(set.Contains(1, 2, 6), false)

	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 9)), false)
	ast.Equal(set.Equal(NewInt16Set(1, 2, 3, 4, 5)), true)
}
