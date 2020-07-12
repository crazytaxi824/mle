// Package slice - compare two slice
package slice

type int64Type struct{}

func Int64() int64Type { return int64Type{} }

// first index of element
func (int64Type) IndexOf(s []int64, element int64) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (int64Type) LastIndexOf(s []int64, element int64) int {
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
func (int64Type) Equal(a, b []int64) bool {
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
func (it int64Type) Contains(a, sub []int64) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it int64Type) ContainsAny(a, sub []int64) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, return a new slice, do not affect the original slice
func (int64Type) Insert(s []int64, element int64, index int) []int64 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		index = lenS
	}

	result := make([]int64, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, return a new slice, do not affect the original slice
func (int64Type) DeleteByIndex(s []int64, index int) []int64 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]int64, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, return a new slice, do not affect the original slice
func (it int64Type) DeleteN(s []int64, element int64, n int) []int64 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
