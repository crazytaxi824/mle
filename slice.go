// Package mle - make life easier
package mle

func IndexOfInt(s []int, n int) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func IndexOfInt64(s []int64, n int64) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func IndexOfInt32(s []int32, n int32) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func IndexOfInt16(s []int16, n int16) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func IndexOfInt8(s []int8, n int8) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}
