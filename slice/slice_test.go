package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var myNum = []int16{1, 2, 3, 4, 5, 4, 3, 2, 1}

func TestIndexOfInt(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(Int16().IndexOf(myNum, 2), 1)
	ast.Equal(Int16().LastIndexOf(myNum, 2), 7)

	ast.Equal(Int16().Equal(myNum, []int16{1, 2, 3, 4, 5, 4, 3, 2, 1}), true)
	ast.Equal(Int16().Equal(myNum, []int16{1, 2, 3, 4, 5, 4}), false)

	ast.Equal(Int16().SameElements(myNum, []int16{1, 1, 2, 2, 3, 3, 4, 4, 5}), true)
	ast.Equal(Int16().SameElements(myNum, []int16{1, 1, 2, 2, 3, 3, 3, 4, 5}), false)

	// ast.Equal(Int().Contains(myNum, []int{1, 2, 5}), true)
	// ast.Equal(Int().Contains(myNum, []int{1, 2, 3, 7}), false)
	//
	// ast.Equal(Int().SameElements(Int().Intersection(myNum, []int{1, 2, 3}), []int{1, 2, 3}), true)
	// ast.Equal(Int().SameElements(Int().Union(myNum, []int{7, 8, 9}), []int{1, 2, 3, 4, 5, 7, 8, 9}), true)
	//
	// ast.Equal(Int().SameElements(Int().Difference(myNum, []int{1, 2, 3}), []int{4, 5}), true)
	//
	// ast.Equal(Int().SameElements(Int().SymmetricDifference(myNum, []int{1, 2, 7, 8}), []int{3, 4, 5, 7, 8}), true)
	// ast.Equal(Int().SameElements(Int().SymmetricDifference(myNum, []int{1, 2, 3, 4, 5}), []int{}), true)
	//
	// ast.Equal(bytes.ContainsAny([]byte("abcdefg"), "gay"), true)
}
