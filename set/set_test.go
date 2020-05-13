package set

import (
	"testing"
)

var s = []int16{1, 2, 3, 4, 5}

func TestSet(t *testing.T) {
	var set = NewFrom(s)
	t.Log(set.Len())

	t.Log(set.Pop())
	t.Log(set.Len())
}
