package example

import (
	"crypto/rand"
	"math/big"
	"sort"
	"testing"

	"github.com/crazytaxi824/mle/tree/avltree"
)

func Benchmark_AVLTreeAdd(b *testing.B) {
	result := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		result = append(result, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := avltree.NewTree()
		for _, v := range result {
			tree.Insert(v, nil)
		}
	}
	b.ReportAllocs()
}

func Benchmark_SliceAdd(b *testing.B) {
	result := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		result = append(result, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int64, 0, 10000)
		s = append(s, result...)
		_ = s
	}
	b.ReportAllocs()
}

func Benchmark_MapAdd(b *testing.B) {
	result := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		result = append(result, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int64]struct{}, 10000)
		for _, v := range result {
			m[v] = struct{}{}
		}
	}
	b.ReportAllocs()
}

func Benchmark_TreeRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := avltree.NewTree()
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

func Benchmark_TreeSearch(b *testing.B) {
	tree := avltree.NewTree()
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

func Benchmark_SliceSearch(b *testing.B) {
	s := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		s = append(s, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range s {
			if s[i] == 1000 {
				_ = s[i]
			}
		}
	}
	b.ReportAllocs()
}

func Benchmark_MapSearch(b *testing.B) {
	m := make(map[int64]struct{}, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		m[r.Int64()] = struct{}{}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m[1000]
	}
	b.ReportAllocs()
}

func Benchmark_TreeSort(b *testing.B) {
	tree := avltree.NewTree()
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

func Benchmark_SliceSort(b *testing.B) {
	s := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		s = append(s, r.Int64())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}
	b.ReportAllocs()
}
