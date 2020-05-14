package util

import (
	frand "math/rand"
	"time"
)

// 赋值给slice，n 次
func InitSlice(n int) []int64 {
	r := frand.New(frand.NewSource(time.Now().UnixNano()))

	var s = make([]int64, n)
	for i := 0; i < n; i++ {
		s[i] = r.Int63n(100000)
	}
	return s
}

func InitMap(n int) map[int64]struct{} {
	r := frand.New(frand.NewSource(time.Now().UnixNano()))
	m := make(map[int64]struct{}, n)

	for i := 0; i < n; i++ {
		m[r.Int63n(100000)] = struct{}{}
	}

	return m
}
