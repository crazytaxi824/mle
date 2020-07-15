package conv

import (
	"bytes"
	"errors"
	"unicode/utf16"
	"unicode/utf8"
)

var ErrNotEven = errors.New("length of []byte must be an even number")

// UTF16toUTF8 字符转换
func UTF16toUTF8(b []byte) ([]byte, error) {
	if len(b)%2 != 0 {
		return nil, ErrNotEven
	}

	u16s := make([]uint16, 1)
	result := &bytes.Buffer{}
	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		result.Write(b8buf[:n])
	}
	return result.Bytes(), nil
}
