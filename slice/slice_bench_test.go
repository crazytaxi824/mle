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

func BenchmarkInt16Type_ContainsAny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int().ContainsAny(myNum, []int{7, 8, 9, 0, 11, 13, 45})
	}
	b.ReportAllocs()
}

func BenchmarkMapSearch(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < 300; i++ {
		m[i] = struct{}{}
	}

	for i := 0; i < b.N; i++ {
		_ = m[20]
	}
	b.ReportAllocs()
}

func BenchmarkSlicesSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < 2; n++ {
			for m := 0; m < 150; m++ {
				if n == m {

				}
			}
		}
	}
	b.ReportAllocs()
}
