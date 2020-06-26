// Package conv - convert types
package conv

import (
	"reflect"
	"unsafe"
)

// BytesToStr []byte转string
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) // #nosec
}

// StrToByte 字符串转[]byte
func StrToByte(s string) []byte {
	x := (*reflect.StringHeader)(unsafe.Pointer(&s)) // #nosec

	var h reflect.SliceHeader
	h.Data = x.Data
	h.Len = x.Len
	h.Cap = x.Len

	return *(*[]byte)(unsafe.Pointer(&h)) // #nosec
}
