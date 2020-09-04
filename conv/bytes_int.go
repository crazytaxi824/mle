package conv

import (
	"encoding/binary"
	"errors"
)

const (
	uint8Len = 1 << iota
	uint16Len
	uint32Len
	uint64Len
)

var (
	ErrWrongLength1 = errors.New("length of bytes must be 1")
	ErrWrongLength2 = errors.New("length of bytes must be 2")
	ErrWrongLength4 = errors.New("length of bytes must be 4")
	ErrWrongLength8 = errors.New("length of bytes must be 8")
)

// BytesToInt8 byte[1] -> int8
func BytesToInt8(b []byte) (int8, error) {
	l := len(b)
	if l != uint8Len {
		return 0, ErrWrongLength1
	}

	return int8(b[0]), nil
}

func BytesToUint8(b []byte) (uint8, error) {
	l := len(b)
	if l != uint8Len {
		return 0, ErrWrongLength1
	}

	return b[0], nil
}

// BytesToInt16 byte[2] -> int16
func BytesToInt16(b []byte, bigEndian ...bool) (int16, error) {
	l := len(b)
	if l != uint16Len {
		return 0, ErrWrongLength2
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return int16(binary.LittleEndian.Uint16(b)), nil
	}

	// BigEndian
	return int16(binary.BigEndian.Uint16(b)), nil
}

func BytesToUint16(b []byte, bigEndian ...bool) (uint16, error) {
	l := len(b)
	if l != uint16Len {
		return 0, ErrWrongLength2
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return binary.LittleEndian.Uint16(b), nil
	}

	// BigEndian
	return binary.BigEndian.Uint16(b), nil
}

// BytesToInt32 byte[4] -> int32
func BytesToInt32(b []byte, bigEndian ...bool) (int32, error) {
	l := len(b)
	if l != uint32Len {
		return 0, ErrWrongLength4
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return int32(binary.LittleEndian.Uint32(b)), nil
	}

	// BigEndian
	return int32(binary.BigEndian.Uint32(b)), nil
}

func BytesToUint32(b []byte, bigEndian ...bool) (uint32, error) {
	l := len(b)
	if l != uint32Len {
		return 0, ErrWrongLength4
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return binary.LittleEndian.Uint32(b), nil
	}

	// BigEndian
	return binary.BigEndian.Uint32(b), nil
}

// BytesToInt64 byte[8] -> int64
func BytesToInt64(b []byte, bigEndian ...bool) (int64, error) {
	l := len(b)
	if l != uint64Len {
		return 0, ErrWrongLength8
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return int64(binary.LittleEndian.Uint64(b)), nil
	}

	// BigEndian
	return int64(binary.BigEndian.Uint64(b)), nil
}

func BytesToUint64(b []byte, bigEndian ...bool) (uint64, error) {
	l := len(b)
	if l != uint64Len {
		return 0, ErrWrongLength8
	}

	be := true
	if len(bigEndian) != 0 {
		be = bigEndian[0]
	}

	if !be {
		// LittleEndian
		return binary.LittleEndian.Uint64(b), nil
	}

	// BigEndian
	return binary.BigEndian.Uint64(b), nil
}
