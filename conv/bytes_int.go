package conv

import (
	"encoding/binary"
	"errors"
	"strconv"
)

type UintValue uint64

const (
	uint16Len = 1 << 1
	uint32Len = 1 << 2
	uint64Len = 1 << 3
	// byteBit   = 1 << 8
)

// BytesToInt byte -> int
func BytesToInt(b []byte) (UintValue, error) {
	l := len(b)
	switch {
	case l <= uint16Len && l > 0:
		// 给 bytes 补到2位
		if l < uint16Len {
			b = append([]byte{0}, b...)
		}
		return UintValue(binary.BigEndian.Uint16(b)), nil
	case l <= uint32Len:
		// 给 bytes 补到4位
		if l < uint32Len {
			b = append([]byte{0}, b...)
		}
		return UintValue(binary.BigEndian.Uint32(b)), nil
	case l <= uint64Len:
		// 给 bytes 补到8位
		tmp := make([]byte, uint64Len)
		if l < uint64Len {
			for i, j := l-1, uint16Len-1; i >= 0; i, j = i-1, j-1 {
				tmp[i] = b[i]
			}
		}
		return UintValue(binary.BigEndian.Uint64(b)), nil
	}
	return 0, errors.New("length of bytes must greater then 0 < len(bytes) <=8")
}

func (v UintValue) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

func (v UintValue) Int() (int, error) {
	if v >= 1<<63 {
		return 0, errors.New("value overflows int64")
	}
	return int(v), nil
}

func (v UintValue) Int64() (int64, error) {
	if v >= 1<<63 {
		return 0, errors.New("value overflows int64")
	}
	return int64(v), nil
}

func (v UintValue) Int32() (int32, error) {
	if v >= 1<<31 {
		return 0, errors.New("value overflows int32")
	}
	return int32(v), nil
}

func (v UintValue) Int16() (int16, error) {
	if v >= 1<<15 {
		return 0, errors.New("value overflows int16")
	}
	return int16(v), nil
}

func (v UintValue) Int8() (int8, error) {
	if v >= 1<<7 {
		return 0, errors.New("value overflows int8")
	}
	return int8(v), nil
}

func (v UintValue) Uint64() (uint64, error) {
	return uint64(v), nil
}

func (v UintValue) Uint32() (uint32, error) {
	if v >= 1<<32 {
		return 0, errors.New("value overflows uint32")
	}
	return uint32(v), nil
}

func (v UintValue) Uint16() (uint16, error) {
	if v >= 1<<16 {
		return 0, errors.New("value overflows uint16")
	}
	return uint16(v), nil
}

func (v UintValue) Uint8() (uint8, error) {
	if v >= 1<<8 {
		return 0, errors.New("value overflows uint8")
	}
	return uint8(v), nil
}

// Uint64ToBytes Uint64 -> Bytes
func Uint64ToBytes(n uint64) []byte {
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, n)
	return tmp
}

func Uint32ToBytes(n uint32) []byte {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, n)
	return tmp
}

func Uint16ToBytes(n uint16) []byte {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, n)
	return tmp
}
