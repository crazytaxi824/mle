package conv

import (
	"encoding/binary"
	"testing"
)

// b 不能为空
func TestBytesToIntLengthEquals0(t *testing.T) {
	b := make([]byte, 0)
	_, err := BytesToInt(b)
	if err == nil {
		t.Fatal()
	}
}

// b length 不能大于 8
func TestBytesToIntLengthGreaterThan8(t *testing.T) {
	b := make([]byte, 9)
	_, err := BytesToInt(b)
	if err == nil {
		t.Fatal()
	}
}

func TestBytesToInt16(t *testing.T) {
	src := []byte{2}
	dst := []byte{0, 2}

	v, err := BytesToInt(src)
	if err != nil {
		t.Error(err)
		return
	}

	i, err := v.Uint16()
	if err != nil {
		t.Error(err)
		return
	}

	if i != binary.BigEndian.Uint16(dst) {
		t.Fatal()
	}
}

func TestBytesToInt32(t *testing.T) {
	src := []byte{1, 2, 3}
	dst := []byte{0, 1, 2, 3}

	v, err := BytesToInt(src)
	if err != nil {
		t.Error(err)
		return
	}

	i, err := v.Uint32()
	if err != nil {
		t.Error(err)
		return
	}

	if i != binary.BigEndian.Uint32(dst) {
		t.Fatal()
	}
}

func TestBytesToInt64(t *testing.T) {
	src := []byte{1, 2, 3, 4, 5, 6}
	dst := []byte{0, 0, 1, 2, 3, 4, 5, 6}

	v, err := BytesToInt(src)
	if err != nil {
		t.Error(err)
		return
	}

	i, err := v.Uint64()
	if err != nil {
		t.Error(err)
		return
	}

	if i != binary.BigEndian.Uint64(dst) {
		t.Fatal()
	}
}
