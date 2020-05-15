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
func (it stringType) ContainsAny(a, sub []string) bool {
	for k := range sub {
		if it.IndexOf(a, sub[k]) != -1 {
			return true
		}
	}
	return false
}

// insert element in the certain index
func (stringType) Insert(s []string, element string, index int) []string {
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
func (stringType) DeleteByIndex(s []string, index int) []string {
	if index < 0 || index > len(s)-1 {
		return s
	}

	result := s[:index:index]
	result = append(result, s[index+1:]...)
	return result
}

// n <= 0 delete all element
func (it stringType) DeleteN(s []string, element string, n int) []string {
	for i := 0; n <= 0 || i < n; i++ {
		index := it.IndexOf(s, element)
		if index < 0 {
			return s
		}

		s = it.DeleteByIndex(s, index)
	}
	return s
}
