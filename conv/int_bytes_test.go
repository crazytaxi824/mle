package conv

import (
	"reflect"
	"testing"
)

func TestInt16ToBytes(t *testing.T) {
	// 测试正数
	i := int16(256)

	r := Int16ToBytes(i)

	if !reflect.DeepEqual(r, []byte{1, 0}) {
		t.Fail()
	}

	// 测试负数
	i = -1

	r = Int16ToBytes(i)

	if !reflect.DeepEqual(r, []byte{255, 255}) {
		t.Fail()
	}
}

func TestInt32ToBytes(t *testing.T) {
	// 测试正数
	i := int32(256)

	r := Int32ToBytes(i)

	if !reflect.DeepEqual(r, []byte{0, 0, 1, 0}) {
		t.Fail()
	}

	// 测试负数
	i = -1

	r = Int32ToBytes(i)

	if !reflect.DeepEqual(r, []byte{255, 255, 255, 255}) {
		t.Fail()
	}
}

func TestInt64ToBytes(t *testing.T) {
	// 测试正数
	i := int64(256)

	r := Int64ToBytes(i)

	if !reflect.DeepEqual(r, []byte{0, 0, 0, 0, 0, 0, 1, 0}) {
		t.Fail()
	}

	// 测试负数
	i = -1

	r = Int64ToBytes(i)

	if !reflect.DeepEqual(r, []byte{255, 255, 255, 255, 255, 255, 255, 255}) {
		t.Fail()
	}
}

func TestUint16ToBytes(t *testing.T) {
	i := uint16(256)

	r := Uint16ToBytes(i)

	if !reflect.DeepEqual(r, []byte{1, 0}) {
		t.Fail()
	}
}

func TestUint32ToBytes(t *testing.T) {
	i := uint32(256)

	r := Uint32ToBytes(i)

	if !reflect.DeepEqual(r, []byte{0, 0, 1, 0}) {
		t.Fail()
	}
}

func TestUint64ToBytes(t *testing.T) {
	i := uint64(256)

	r := Uint64ToBytes(i)

	if !reflect.DeepEqual(r, []byte{0, 0, 0, 0, 0, 0, 1, 0}) {
		t.Fail()
	}
}
