// Package slice - compare two slice
package slice

type uint16Type struct{}

func Uint16() uint16Type { return uint16Type{} }

// first index of element
func (uint16Type) IndexOf(s []uint16, element uint16) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (uint16Type) LastIndexOf(s []uint16, element uint16) int {
	if len(s) < 1 {
		return -1
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == element {
			return i
		}
	}
	return -1
}

// is A == B ?
func (uint16Type) Equal(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

// is A contains all elements of SUB ?
func (it uint16Type) Contains(a, sub []uint16) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it uint16Type) ContainsAny(a, sub []uint16) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, return a new slice, do not affect the original slice
func (uint16Type) Insert(s []uint16, element uint16, index int) []uint16 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		index = lenS
	}

	result := make([]uint16, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, return a new slice, do not affect the original slice
func (uint16Type) DeleteByIndex(s []uint16, index int) []uint16 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]uint16, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, return a new slice, do not affect the original slice
func (it uint16Type) DeleteN(s []uint16, element uint16, n int) []uint16 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
