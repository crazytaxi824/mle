package avltree

import (
	"sort"
	"testing"
)

// var s = []int{10, 9, 8}
var s = []int{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

// add
func TestAVLTree_Add(t *testing.T) {
	tree := NewAVLTree()
	for _, v := range s {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	PrintAllNode(tree.root)

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
	tree := NewAVLTree()
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

	PrintAllNode(tree.root)

	if tree.Size() != len(s)-5 {
		t.Fail()
	}
}

func TestAVLTree_Delete2(t *testing.T) {
	ss := []int{100, 50, 150, 70, 30, 40, 20, 10, 25}
	tree := NewAVLTree()
	for _, v := range ss {
		err := tree.Add(v, struct{}{})
		if err != nil {
			t.Error(err)
			return
		}
	}

	_ = tree.DeleteFromOrder(10)
	_ = tree.DeleteFromOrder(40)

	PrintAllNode(tree.root)
}

// sort
func TestAVLTree_Sort(t *testing.T) {
	tree := NewAVLTree()
	for _, v := range s {
		err := tree.Add(v, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	PrintAllNode(tree.root)

	for _, v := range tree.Sort() {
		t.Log(v.order)
	}
}

// search
func BenchmarkSearchInAVLTree(b *testing.B) {
	tree := NewAVLTree()
	for i := 0; i < 1000; i++ {
		err := tree.Add(i, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}

	for i := 0; i < b.N; i++ {
		n := tree.Find(900)
		_, ok := n.value.(struct{})
		if !ok {
			b.Fail()
		}
	}
	b.ReportAllocs()
}

func BenchmarkSearchInSlice(b *testing.B) {
	ss := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		ss[i] = i
	}

	for i := 0; i < b.N; i++ {
		for _, v := range ss {
			if v == 900 {

			}
		}
	}
	b.ReportAllocs()
}

// add
func BenchmarkAVLTree_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := NewAVLTree()
		for n := 0; n < 1000; n++ {
			err := tree.Add(n, struct{}{})
			if err != nil {
				b.Error(err)
			}
		}
	}
	b.ReportAllocs()
}

func BenchmarkAppendInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var ss []int
		for n := 0; n < 1000; n++ {
			_ = append(ss, n)
		}
	}
	b.ReportAllocs()
}

// sort
func BenchmarkAVLTree_Sort(b *testing.B) {
	tree := NewAVLTree()
	for _, v := range s {
		err := tree.Add(v, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		tree.Sort()
	}
	b.ReportAllocs()
}

func BenchmarkSortSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
	b.ReportAllocs()
}

func TestDeleteAVLRoot(t *testing.T) {
	tree := NewAVLTree()
	for i := 0; i < 5; i++ {
		tree.Add(i, struct{}{})
	}

	for i := 0; i < 5; i++ {
		t.Log(i)
		err := tree.DeleteFromOrder(i)
		if err != nil {
			t.Error(err)
			return
		}
		PrintAllNode(tree.root)
	}
}
