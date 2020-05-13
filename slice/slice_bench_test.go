package slice

import (
	"testing"
)

func BenchmarkSameElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int().SameElements(myNum, []int{1, 1, 2, 2, 3, 3, 4, 4, 5})
	}
	b.ReportAllocs()
}

func BenchmarkDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int().Difference(myNum, []int{1, 2, 3})
	}
	b.ReportAllocs()
}
