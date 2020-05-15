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
func (it runeType) ContainsAny(a, sub []rune) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index
func (runeType) Insert(s []rune, element rune, index int) []rune {
	if index < 0 {
		index = 0
	}

	lenS := len(s)
	if index >= lenS {
		s = append(s, element)
		return s
	}

	result := s[:index:index]
	result = append(result, element)
	result = append(result, s[index:]...)

	return result
}

// delete by index
func (runeType) DeleteByIndex(s []rune, index int) []rune {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := s[:index:index]
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element
func (it runeType) DeleteN(s []rune, element rune, n int) []rune {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
