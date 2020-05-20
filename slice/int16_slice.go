// Package slice - compare two slice
package slice

type int16Type struct{}

func Int16() int16Type { return int16Type{} }

// first index of element
func (int16Type) IndexOf(s []int16, element int16) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (int16Type) LastIndexOf(s []int16, element int16) int {
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
func (int16Type) Equal(a, b []int16) bool {
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
func (it int16Type) Contains(a, sub []int16) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (it int16Type) ContainsAny(a, sub []int16) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index, do not affect the original slice
func (int16Type) Insert(s []int16, element int16, index int) []int16 {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		return append(s, element)
	}

	result := make([]int16, index, lenS+1)
	copy(result, s)
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index, do not affect the original slice
func (int16Type) DeleteByIndex(s []int16, index int) []int16 {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := make([]int16, index, len(s)-1)
	copy(result, s)
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element, do not affect the original slice
func (it int16Type) DeleteN(s []int16, element int16, n int) []int16 {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
