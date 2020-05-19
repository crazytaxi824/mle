package tree

import (
	"testing"

	"github.com/crazytaxi824/mle/tree/avltree"
)

// var s = []int{10, 9, 8}
var s = []int{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

// add
func TestAVLTree_Add(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range s {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	avltree.PrintAllNode(tree.Root())

	n := tree.Find(82)
	if n == nil {
		t.Fail()
	}

	t.Log(tree.Root().Order())
	t.Log(tree.Smallest().Order())
	t.Log(tree.Biggest().Order())

	if tree.Size() != len(s) {
		t.Fail()
	}
}

// delete
func TestAVLTree_Delete(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range s {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	_ = tree.DeleteFromOrder(84)
	_ = tree.DeleteFromOrder(85)
	_ = tree.DeleteFromOrder(120)
	_ = tree.DeleteFromOrder(87)
	_ = tree.DeleteFromOrder(86)

	avltree.PrintAllNode(tree.Root())

	if tree.Size() != len(s)-5 {
		t.Fail()
	}
}

func TestAVLTree_Delete2(t *testing.T) {
	ss := []int{100, 50, 150, 70, 30, 40, 20, 10, 25}
	tree := avltree.NewTree()
	for _, v := range ss {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	_ = tree.DeleteFromOrder(10)
	_ = tree.DeleteFromOrder(40)

	avltree.PrintAllNode(tree.Root())
}

func TestDeleteAVLRoot(t *testing.T) {
	tree := avltree.NewTree()
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
func TestAVLTree_Sort(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range s {
		err := tree.Add(v, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	avltree.PrintAllNode(tree.Root())

	for _, v := range tree.Sort() {
		t.Log(v.Order())
	}
}
