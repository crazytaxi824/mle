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

var del = []int{50, 20, 65, 15, 35, 55, 70, 68, 80, 90}

func TestDelete(t *testing.T) {
	tree := NewRBTree()
	for _, v := range del {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	err := tree.DeleteFromOrder(55)
	if err != nil {
		t.Error(err)
		return
	}

	PrintAllNode(tree.root)
}
