package example

import (
	"local/src/rbtree"
	"testing"
)

var items = []int64{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70, 65}

func Test_Insert(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_Delete(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}

	tree.Delete(30)
	tree.Delete(15)
	tree.Delete(40)

	PrintALL(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_Search(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Insert(v, v*100)
	}

	n := tree.Search(30)
	if n == nil {
		t.Error("30 not exist")
		return
	}

	t.Log(n.Index())
	t.Log(n.Value())

	tree.Delete(30)

	n = tree.Search(30)
	if n == nil {
		t.Log("30 not exist")
		return
	}
}

func Test_SmallAndLarge(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Insert(v, v*100)
	}

	t.Log(tree.Smallest().Index(), tree.Largest().Index())
}

func Test_Sort(t *testing.T) {
	tree := rbtree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	if err := CheckAllNodes(tree); err != nil {
		t.Error(err)
	}

	var ns []int64
	for _, v := range tree.Sort() {
		ns = append(ns, v.Index())
	}
	t.Log(ns)
}
