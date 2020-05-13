// Package conv - convert types
package conv

import (
	"unsafe"
)

// BytesToStr []byte转string
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // #nosec
}

// StrToByte 字符串转[]byte
func StrToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s)) // #nosec
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h)) // #nosec
}
