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
	if n == nil {
		t.Fail()
	}
}

func BenchmarkSlice(b *testing.B) {
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

func BenchmarkAVLTree(b *testing.B) {
	tree := NewAVLTree()
	for i := 0; i < 1000; i++ {
		err := tree.Add(i, i)
		if err != nil {
			b.Error(err)
			return
		}
	}

	for i := 0; i < b.N; i++ {
		n := tree.Find(900)
		_, ok := n.value.(int)
		if !ok {
			b.Fail()
		}
	}
	b.ReportAllocs()
}

func TestAVLTree_Delete(t *testing.T) {
	tree := NewAVLTree()
	for _, v := range s {
		err := tree.Add(v, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	err := tree.Delete(84)
	if err != nil {
		t.Error(err)
		return
	}

	PrintAllNode(tree.root)
}
