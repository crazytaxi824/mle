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
func (intType) ContainsAny(a, sub []int) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
