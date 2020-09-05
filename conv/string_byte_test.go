package conv

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStringToBytes(t *testing.T) {
	s := "abc"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))

	t.Logf("%+v", sh)
	t.Log(*(*uint8)(unsafe.Pointer(sh.Data)))
	t.Log(*(*uint8)(unsafe.Pointer(sh.Data + 1)))
	t.Log(*(*uint8)(unsafe.Pointer(sh.Data + 2)))

	var bh reflect.SliceHeader
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len

	b := *(*[]byte)(unsafe.Pointer(&bh))
	t.Log(b)

	// string 本质也是一个结构体，而且是 const 不可变类型
	// 所以这里如果操作 byte[0] 写入新的值，会报错 fatal error
	// b[1] = 100
}

func TestBytesToStr(t *testing.T) {
	b := []byte{97, 98, 99}
	s := BytesToStr(b)
	if s != "abc" {
		t.Fail()
	}

	// 更改原始 bytes 会影响 string 的值
	b[0] = 100
	if s != "dbc" {
		t.Fail()
	}
}
