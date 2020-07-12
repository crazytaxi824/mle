// Package slice - compare two slice
package slice

type intType struct{}

func Int() intType { return intType{} }

// first index of element
func (intType) IndexOf(s []int, element int) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (intType) LastIndexOf(s []int, element int) int {
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
func (intType) Equal(a, b []int) bool {
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
func (it intType) Contains(a, sub []int) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it intType) ContainsAny(a, sub []int) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, return a new slice, do not affect the original slice
func (intType) Insert(s []int, element, index int) []int {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		index = lenS
	}

	result := make([]int, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, return a new slice, do not affect the original slice
func (intType) DeleteByIndex(s []int, index int) []int {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]int, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, return a new slice, do not affect the original slice
func (it intType) DeleteN(s []int, element, n int) []int {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
