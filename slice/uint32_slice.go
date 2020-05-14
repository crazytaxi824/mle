// Package slice - compare two slice
package slice

type uint32Type struct{}

func Uint32() uint32Type { return uint32Type{} }

// first index of element
func (uint32Type) IndexOf(s []uint32, element uint32) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (uint32Type) LastIndexOf(s []uint32, element uint32) int {
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
func (uint32Type) Equal(a, b []uint32) bool {
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
func (it uint32Type) Contains(a, sub []uint32) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it uint32Type) ContainsAny(a, sub []uint32) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index
func (uint32Type) Insert(s []uint32, element uint32, index int) []uint32 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		s = append(s, element)
		return s
	}

	result := s[:index:index]
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index
func (uint32Type) DeleteByIndex(s []uint32, index int) []uint32 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := s[:index:index]
	result = append(result, s[index+1:]...)
	return result
}

// n < 0 delete all element
func (it uint32Type) DeleteN(s []uint32, element uint32, n int) []uint32 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
