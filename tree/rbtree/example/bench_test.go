package example

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/crazytaxi824/mle/tree/rbtree"
)

func Benchmark_RBTreeAdd(b *testing.B) {
	result := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		result = append(result, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := rbtree.NewTree()
		for _, v := range result {
			tree.Insert(v, nil)
		}
	}
	b.ReportAllocs()
}

func Benchmark_RBTreeRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := rbtree.NewTree()
		rms := make([]int64, 10)
		for i := 0; i < 10000; i++ {
			r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
			tree.Insert(r.Int64(), nil)

			if i%100 == 0 {
				rms = append(rms, r.Int64())
			}
		}

		for _, v := range rms {
			tree.Delete(v)
		}
	}
	b.ReportAllocs()
}

func Benchmark_RBTreeSearch(b *testing.B) {
	tree := rbtree.NewTree()
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		tree.Insert(r.Int64(), nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(1000)
	}
	b.ReportAllocs()
}

func Benchmark_RBTreeSort(b *testing.B) {
	tree := rbtree.NewTree()
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		tree.Insert(r.Int64(), nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Sort()
	}
	b.ReportAllocs()
}
