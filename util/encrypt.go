package util

import (
	"crypto/md5"  // #nosec, not recommended
	"crypto/sha1" // #nosec, not recommended
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

const (
	SHA256 = iota
	SHA512
	SHA1 // not recommended
	MD5  // not recommended
)

// hash 加密
func EncryptHash(msg, salt []byte, flag int) ([]byte, error) {
	var h hash.Hash
	switch flag {
	case SHA256:
		h = sha256.New()
	case SHA512:
		h = sha512.New()
	case SHA1:
		h = sha1.New() // #nosec, not recommended
	case MD5:
		h = md5.New() // #nosec, not recommended
	default:
		h = sha256.New()
	}

	_, err := h.Write(msg)
	if err != nil {
		return nil, err
	}
	return h.Sum(salt), nil
}
