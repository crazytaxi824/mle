package conv

import (
	"encoding/binary"
)

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

func Int64ToBytes(n int64) []byte {
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, uint64(n))
	return tmp
}

func Int32ToBytes(n int32) []byte {
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint32(tmp, uint32(n))
	return tmp
}

func Int16ToBytes(n int16) []byte {
	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, uint16(n))
	return tmp
}
