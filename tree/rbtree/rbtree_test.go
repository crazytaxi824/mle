package rbtree

import (
	"testing"
)

var src = []int{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70}

func TestAdd(t *testing.T) {
	tree := NewRBTree()
	for _, v := range src {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	PrintAllNode(tree.root)
}
