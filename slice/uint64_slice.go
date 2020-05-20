// Package slice - compare two slice
package slice

type uint64Type struct{}

func Uint64() uint64Type { return uint64Type{} }

// first index of element
func (uint64Type) IndexOf(s []uint64, element uint64) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (uint64Type) LastIndexOf(s []uint64, element uint64) int {
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
func (uint64Type) Equal(a, b []uint64) bool {
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
func (it uint64Type) Contains(a, sub []uint64) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it uint64Type) ContainsAny(a, sub []uint64) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, do not affect the original slice
func (uint64Type) Insert(s []uint64, element uint64, index int) []uint64 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		return append(s, element)
	}

	result := make([]uint64, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, do not affect the original slice
func (uint64Type) DeleteByIndex(s []uint64, index int) []uint64 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]uint64, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, do not affect the original slice
func (it uint64Type) DeleteN(s []uint64, element uint64, n int) []uint64 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
