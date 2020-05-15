package tree

import (
	"testing"
)

var s = []int{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

// var s = []int{10, 9, 8}

func TestAVLTree_Add(t *testing.T) {
	tree := NewAVLTree()
	for _, v := range s {
		err := tree.Add(v, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	PrintAllNode(tree.root)

	n := tree.Find(82)
	if n != nil {
		t.Log(n.Order(), n.Value())
	} else {
		t.Log("nil")
	}
}
