// Package slice - compare two slice
package slice

type stringType struct{}

func Strings() stringType { return stringType{} }

// first index of element
func (stringType) IndexOf(s []string, element string) int {
	for k := range s {
		if s[k] == element {
			return k
		}
	}
	return -1
}

// last index of element
func (stringType) LastIndexOf(s []string, element string) int {
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
func (stringType) Equal(a, b []string) bool {
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
func (it stringType) Contains(a, sub []string) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) == -1 {
			return false
		}
	}
	return true
}

// is A contains any element of SUB ?
func (stringType) ContainsAny(a, sub []string) bool {
	for ka := range a {
		for kb := range sub {
			if a[ka] == sub[kb] {
				return true
			}
		}
	}
	return false
}
