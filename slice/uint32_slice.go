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
func (uint32Type) ContainsAny(a, sub []uint32) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
