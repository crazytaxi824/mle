package util

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	frand "math/rand"
	"strconv"
	"time"
)

func TrueRandomIntn(n int64) (int64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		return 0, err
	}

	return r.Int64(), nil
}

func NewFakeRand() *frand.Rand {
	return frand.New(frand.NewSource(time.Now().UnixNano()))
}

// 返回长度为 n 的 string
func RandString(fr *frand.Rand) string {
	s := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(fr.Int63n(1<<63-1), 16)))
	return s[:len(s)-2]
}
