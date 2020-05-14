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
func (uint64Type) ContainsAny(a, sub []uint64) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
