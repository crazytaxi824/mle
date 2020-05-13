package mle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfInt(t *testing.T) {
	var s = []int{1, 2, 3, 4, 5}
	ast := assert.New(t)
	ast.Equal(IndexOfInt(s, 3), 2)
	ast.Equal(IndexOfInt(s, 0), -1)
	ast.Equal(IndexOfInt(s, 1), 0)
}
