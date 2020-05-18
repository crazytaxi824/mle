package tree

import (
	"testing"

	"github.com/crazytaxi824/mle/tree/avltree"
	"github.com/crazytaxi824/mle/tree/rbtree"
)

// ADD
func BenchmarkSlice_Append(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for i := 0; i < 1000; i++ {
			_ = append(s, i)
		}
	}
	b.ReportAllocs()
}

func BenchmarkAVLTree_ADD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		avl := avltree.NewTree()
		for i := 0; i < 1000; i++ {
			err := avl.Add(i, struct{}{})
			if err != nil {
				b.Error(err)
				return
			}
		}
	}
	b.ReportAllocs()
}

func BenchmarkRBTree_ADD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rb := rbtree.NewTree()
		for i := 0; i < 1000; i++ {
			err := rb.Add(i, struct{}{})
			if err != nil {
				b.Error(err)
				return
			}
		}
	}
	b.ReportAllocs()
}

// FIND
func BenchmarkSlice_Find(b *testing.B) {
	s := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = i
	}
	for i := 0; i < b.N; i++ {
		for _, v := range s {
			if v == 500 {
			}
		}
	}
	b.ReportAllocs()
}

func BenchmarkAVLTree_Find(b *testing.B) {
	avl := avltree.NewTree()
	for i := 0; i < 1000; i++ {
		err := avl.Add(i, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		avl.Find(500)
	}
	b.ReportAllocs()
}

func BenchmarkRBTree_Find(b *testing.B) {
	rb := rbtree.NewTree()
	for i := 0; i < 1000; i++ {
		err := rb.Add(i, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		rb.Find(500)
	}
	b.ReportAllocs()
}

// DELETE
const size = 1000000

func BenchmarkAVLTree_Delete(b *testing.B) {
	avl := avltree.NewTree()
	for i := 0; i < size; i++ {
		err := avl.Add(i, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.StartTimer()
	for i := 0; i < size; i++ {
		err := avl.DeleteFromOrder(i)
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}

func BenchmarkRBTree_Delete(b *testing.B) {
	rb := rbtree.NewTree()
	for i := 0; i < size; i++ {
		err := rb.Add(i, struct{}{})
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.StartTimer()
	for i := 0; i < size; i++ {
		err := rb.DeleteFromOrder(i)
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}

// sort
var s = []int{100, 150, 50, 30, 70, 120, 20, 10, 40, 80, 81, 82, 83, 84, 85, 86, 87, 88}

func BenchmarkAVLTree_Sort(b *testing.B) {
	tree := avltree.NewTree()
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

func BenchmarkRBTree_Sort(b *testing.B) {
	tree := rbtree.NewTree()
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
