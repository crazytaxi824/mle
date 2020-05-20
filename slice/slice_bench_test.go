package slice

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		r := src[:3:3]
		r = append(r, 100)
	}
	b.ReportAllocs()
}

func BenchmarkCopy(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		r := make([]int, 3, 4)
		copy(r, src)
		r = append(r, 100)
	}
	b.ReportAllocs()
}
