package example

import (
	"testing"

	"github.com/crazytaxi824/mle/tree/avltree"
)

var items = []int64{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

func Test_Add(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	PrintTree(tree)
	if re := CheckAllNodes(tree); re != nil {
		t.Error(re)
	}
}

func Test_Remove(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	PrintTree(tree)
	if err := CheckAllNodes(tree); err != nil {
		t.Error(err)
	}

	// remove item 的时候 depth 计算问题
	tree.Delete(84)
	tree.Delete(85)
	tree.Delete(120)
	tree.Delete(87)
	tree.Delete(86)

	PrintTree(tree)
	if err := CheckAllNodes(tree); err != nil {
		t.Error(err)
	}
}

func Test_Sort(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range items {
		tree.Insert(v, nil)
	}

	var ns []int64
	for _, v := range tree.Sort() {
		ns = append(ns, v.Index())
	}
	t.Log(ns)
}

func Test_Search(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range items {
		tree.Insert(v, v*100)
	}

	t.Log(tree.Search(80).Index(), tree.Search(80).Value())
}

func Test_SmallAndLarge(t *testing.T) {
	tree := avltree.NewTree()
	for _, v := range items {
		tree.Insert(v, v*100)
	}

	t.Log(tree.Smallest().Index(), tree.Largest().Index())
}
