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

	ast.Equal(Int64().Insert(myNum, 100, 0), []int64{100, 1, 2, 3, 4, 5, 4, 3, 2, 1})
}

func TestInsert(t *testing.T) {
	t.Log(Int64().Insert(myNum, 100, 0))
	t.Log(Int64().Insert(myNum, 100, 100))
	t.Log(Int64().Insert(myNum, 100, 3))
}

func TestDelete(t *testing.T) {
	t.Log(Int64().DeleteByIndex(myNum, 8), myNum)
	t.Log(Int64().DeleteByIndex(myNum, 0), myNum)
}

func TestDeleteALl(t *testing.T) {
	t.Log(Int64().DeleteN(myNum, 8, 0), myNum)
	t.Log(Int64().DeleteN(myNum, 1, 0), myNum)

	t.Log(Int64().DeleteN(myNum, 1, 1), myNum)
}
