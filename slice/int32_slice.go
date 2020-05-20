// Package slice - compare two slice
package slice

type int32Type struct{}

func Int32() int32Type { return int32Type{} }

// first index of element
func (int32Type) IndexOf(s []int32, element int32) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (int32Type) LastIndexOf(s []int32, element int32) int {
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
func (int32Type) Equal(a, b []int32) bool {
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
func (it int32Type) Contains(a, sub []int32) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it int32Type) ContainsAny(a, sub []int32) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, do not affect the original slice
func (int32Type) Insert(s []int32, element int32, index int) []int32 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		return append(s, element)
	}

	result := make([]int32, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, do not affect the original slice
func (int32Type) DeleteByIndex(s []int32, index int) []int32 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]int32, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, do not affect the original slice
func (it int32Type) DeleteN(s []int32, element int32, n int) []int32 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
