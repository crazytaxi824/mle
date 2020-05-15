package set

import (
	"container/list"
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
		set.Contains(1)
	}
	b.ReportAllocs()
}

func BenchmarkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var l = list.New()
		for n := 0; n < 100; n++ {
			l.PushBack(n)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = make([]int, 100)
		for n := 0; n < 100; n++ {
			s[n] = n
		}
	}
	b.ReportAllocs()
}
