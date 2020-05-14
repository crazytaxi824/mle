package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var myNum = []int64{1, 2, 3, 4, 5, 4, 3, 2, 1}

func TestSlice(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(Int64().IndexOf(myNum, 2), 1)
	ast.Equal(Int64().LastIndexOf(myNum, 2), 7)

	ast.Equal(Int64().Equal(myNum, []int64{1, 2, 3, 4, 5, 4, 3, 2, 1}), true)
	ast.Equal(Int64().Equal(myNum, []int64{1, 2, 3, 4, 5, 4}), false)

	ast.Equal(Int64().ContainsAny(myNum, []int64{1, 2, 5}), true)
	ast.Equal(Int64().ContainsAny(myNum, []int64{9, 8, 7}), false)
}
