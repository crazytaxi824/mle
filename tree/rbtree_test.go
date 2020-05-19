package tree

import (
	"testing"

	"github.com/crazytaxi824/mle/tree/rbtree"
)

var src = []int{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70}

// add
func TestAdd(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range src {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	rbtree.PrintAllNode(tree.Root())
}

var del = []int{50, 20, 65, 15, 35, 55, 70, 68, 80, 90}

// delete
func TestDelete(t *testing.T) {
	tree := rbtree.NewTree()
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

	rbtree.PrintAllNode(tree.Root())
}

func TestDeleteRBRoot(t *testing.T) {
	tree := rbtree.NewTree()
	for i := 0; i < 5; i++ {
		tree.Add(i, struct{}{})
	}
	t.Log(tree.Size())

	for i := 0; i < 5; i++ {
		err := tree.DeleteFromOrder(i)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(tree.Size())
	}
}

// sort
func TestRbTree_Sort(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range src {
		err := tree.Add(v, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	rbtree.PrintAllNode(tree.Root())

	for _, v := range tree.Sort() {
		t.Log(v.Order())
	}
}