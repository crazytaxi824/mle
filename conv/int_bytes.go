package conv

import (
	"encoding/binary"
)

// Uint64ToBytes Uint64 -> Bytes
func Uint64ToBytes(n uint64, bigEndian ...bool) []byte {
	tmp := make([]byte, 8)

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		binary.LittleEndian.PutUint64(tmp, n)
	} else {
		binary.BigEndian.PutUint64(tmp, n)
	}

	return tmp
}

func Uint32ToBytes(n uint32, bigEndian ...bool) []byte {
	tmp := make([]byte, 4)

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		binary.LittleEndian.PutUint32(tmp, n)
	} else {
		binary.BigEndian.PutUint32(tmp, n)
	}

	return tmp
}

func Uint16ToBytes(n uint16, bigEndian ...bool) []byte {
	tmp := make([]byte, 2)

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		binary.LittleEndian.PutUint16(tmp, n)
	} else {
		binary.BigEndian.PutUint16(tmp, n)
	}

	return tmp
}

func Int64ToBytes(n int64, bigEndian ...bool) []byte {
	tmp := make([]byte, 8)
	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}
	if !be {
		binary.LittleEndian.PutUint64(tmp, uint64(n))
	} else {
		binary.BigEndian.PutUint64(tmp, uint64(n))
	}
	return tmp
}

func Int32ToBytes(n int32, bigEndian ...bool) []byte {
	tmp := make([]byte, 4)
	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}
	if !be {
		binary.LittleEndian.PutUint32(tmp, uint32(n))
	} else {
		binary.BigEndian.PutUint32(tmp, uint32(n))
	}
	return tmp
}

func Int16ToBytes(n int16, bigEndian ...bool) []byte {
	tmp := make([]byte, 2)
	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}
	if !be {
		binary.LittleEndian.PutUint16(tmp, uint16(n))
	} else {
		binary.BigEndian.PutUint16(tmp, uint16(n))
	}
	return tmp
}
