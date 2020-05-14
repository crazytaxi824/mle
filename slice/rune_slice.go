// Package slice - compare two slice
package slice

type runeType struct{}

func Rune() runeType { return runeType{} }

// first index of element
func (runeType) IndexOf(s []rune, element rune) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (runeType) LastIndexOf(s []rune, element rune) int {
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
func (runeType) Equal(a, b []rune) bool {
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
func (it runeType) Contains(a, sub []rune) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (runeType) ContainsAny(a, sub []rune) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
