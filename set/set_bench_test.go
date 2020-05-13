package set

import (
	"testing"
)

func BenchmarkHashSet_ToSlice(b *testing.B) {
	set := NewInt16Set(1, 2, 3, 4, 5)
	for i := 0; i < b.N; i++ {
		set.ToSlice()
	}
	b.ReportAllocs()
}

func BenchmarkHashSetInt16_Contains(b *testing.B) {
	set := NewInt16Set(1, 2, 3, 4, 5)
	for i := 0; i < b.N; i++ {
		set.Contains(1, 2, 3)
	}
	b.ReportAllocs()
}
