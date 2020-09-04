package conv

import (
	"encoding/binary"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToInt8(t *testing.T) {
	type testInt8 struct {
		arg    []byte
		result int8
	}

	ts := []testInt8{
		{[]byte{0b00000001}, 1},
		{[]byte{0b10000000}, -128},
		{[]byte{0b10000001}, -127},
		{[]byte{0b11111111}, -1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		r, err := BytesToInt8(v.arg)
		if err != nil {
			t.Error(err)
			return
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToInt8([]byte{1, 2})
	if !errors.Is(err, ErrWrongLength1) {
		t.Fail()
	}
}

func TestBytesToUint8(t *testing.T) {
	type testUint8 struct {
		arg    []byte
		result uint8
	}

	ts := []testUint8{
		{[]byte{0b00000001}, 1},
		{[]byte{0b10000000}, 1 << 7},
		{[]byte{0b10000001}, 1<<7 + 1},
		{[]byte{0b11111111}, 1<<8 - 1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		r, err := BytesToUint8(v.arg)
		if err != nil {
			t.Error(err)
			return
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToUint8([]byte{1, 2})
	if !errors.Is(err, ErrWrongLength1) {
		t.Fail()
	}
}

func TestBytesToInt16(t *testing.T) {
	type testInt16 struct {
		arg       []byte
		bigEndian bool
		result    int16
	}

	ts := []testInt16{
		{[]byte{0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000}, true, -1 << 15},
		{[]byte{0b10000000, 0b00000001}, true, -1<<15 + 1},
		{[]byte{0b11111111, 0b11111111}, true, -1},

		{[]byte{0b00000001, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b10000000}, false, -1 << 15},
		{[]byte{0b00000001, 0b10000000}, false, -1<<15 + 1},
		{[]byte{0b11111111, 0b11111111}, false, -1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   int16
			err error
		)

		if v.bigEndian {
			r, err = BytesToInt16(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToInt16(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToInt16([]byte{1})
	if !errors.Is(err, ErrWrongLength2) {
		t.Fail()
	}
}

func TestBytesToUint16(t *testing.T) {
	type testUint16 struct {
		arg       []byte
		bigEndian bool
		result    uint16
	}

	ts := []testUint16{
		{[]byte{0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000}, true, 1 << 15},
		{[]byte{0b10000000, 0b00000001}, true, 1<<15 + 1},
		{[]byte{0b11111111, 0b11111111}, true, 1<<16 - 1},

		{[]byte{0b00000001, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b10000000}, false, 1 << 15},
		{[]byte{0b00000001, 0b10000000}, false, 1<<15 + 1},
		{[]byte{0b11111111, 0b11111111}, false, 1<<16 - 1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   uint16
			err error
		)

		if v.bigEndian {
			r, err = BytesToUint16(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToUint16(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToUint16([]byte{1})
	if !errors.Is(err, ErrWrongLength2) {
		t.Fail()
	}
}

func TestBytesToInt32(t *testing.T) {
	type testInt32 struct {
		arg       []byte
		bigEndian bool
		result    int32
	}

	ts := []testInt32{
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000}, true, -1 << 31},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000001}, true, -1<<31 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111}, true, -1},

		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, -1 << 31},
		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b10000000}, false, -1<<31 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111}, false, -1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   int32
			err error
		)

		if v.bigEndian {
			r, err = BytesToInt32(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToInt32(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToInt32([]byte{1})
	if !errors.Is(err, ErrWrongLength4) {
		t.Fail()
	}
}

func TestBytesToUint32(t *testing.T) {
	type testUint32 struct {
		arg       []byte
		bigEndian bool
		result    uint32
	}

	ts := []testUint32{
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000}, true, 1 << 31},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000001}, true, 1<<31 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111}, true, 1<<32 - 1},

		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, 1 << 31},
		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b10000000}, false, 1<<31 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111}, false, 1<<32 - 1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   uint32
			err error
		)

		if v.bigEndian {
			r, err = BytesToUint32(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToUint32(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToUint32([]byte{1})
	if !errors.Is(err, ErrWrongLength4) {
		t.Fail()
	}
}

func TestBytesToInt64(t *testing.T) {
	type testInt64 struct {
		arg       []byte
		bigEndian bool
		result    int64
	}

	ts := []testInt64{
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, true, -1 << 63},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, -1<<63 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111}, true, -1},

		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, -1 << 63},
		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, -1<<63 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111}, false, -1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   int64
			err error
		)

		if v.bigEndian {
			r, err = BytesToInt64(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToInt64(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToInt64([]byte{1})
	if !errors.Is(err, ErrWrongLength8) {
		t.Fail()
	}
}

func TestBytesToUint64(t *testing.T) {
	type testUint64 struct {
		arg       []byte
		bigEndian bool
		result    uint64
	}

	ts := []testUint64{
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, 1},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, true, 1 << 63},
		{[]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000001}, true, 1<<63 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111}, true, 1<<64 - 1},

		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, false, 1},
		{[]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, 1 << 63},
		{[]byte{0b00000001, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b10000000}, false, 1<<63 + 1},
		{[]byte{0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111, 0b11111111}, false, 1<<64 - 1},
	}

	ast := assert.New(t)
	for _, v := range ts {
		var (
			r   uint64
			err error
		)

		if v.bigEndian {
			r, err = BytesToUint64(v.arg)
			if err != nil {
				t.Error(err)
				return
			}
		} else {
			r, err = BytesToUint64(v.arg, v.bigEndian)
			if err != nil {
				t.Error(err)
				return
			}
		}

		if !ast.Equal(r, v.result) {
			t.Fail()
		}
	}

	_, err := BytesToUint64([]byte{1})
	if !errors.Is(err, ErrWrongLength8) {
		t.Fail()
	}
}

func TestConvert(t *testing.T) {
	b := int8(-128)

	c := int16(b)
	t.Log(c)

	r := make([]byte, 2)
	binary.BigEndian.PutUint16(r, uint16(c))
	t.Log(r)

	for _, v := range r {
		t.Log(strconv.FormatInt(int64(v), 2))
	}
}
