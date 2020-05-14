package util

import (
	frand "math/rand"
	"time"
)

// 赋值给slice，n 次
func initSlice(n int) []int {
	r := frand.New(frand.NewSource(time.Now().UnixNano()))

	var s = make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Intn(100000)
	}
	return s
}

func initMap(n int) map[int]struct{} {
	r := frand.New(frand.NewSource(time.Now().UnixNano()))
	m := make(map[int]struct{}, n)

	for i := 0; i < n; i++ {
		m[r.Intn(100000)] = struct{}{}
	}

	return m
}
