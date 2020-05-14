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
func (uint16Type) ContainsAny(a, sub []uint16) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
