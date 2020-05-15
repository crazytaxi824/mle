package util

import (
	"crypto/rand"
	"math/big"
	frand "math/rand"
	"testing"
	"time"
)

func BenchmarkTrueRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(1000))
		r.Int64()
	}
	b.ReportAllocs()
}

func BenchmarkFakeRandomSameSeed(b *testing.B) {
	r := frand.New(frand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		r.Int63n(1000)
	}
	b.ReportAllocs()
}

func BenchmarkFakeRandomNewSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		frand.New(frand.NewSource(time.Now().UnixNano())).Int63n(1000)
	}
	b.ReportAllocs()
}
