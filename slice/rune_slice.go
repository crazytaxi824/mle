package slice

type runeType struct{}

func Rune() runeType { return runeType{} }

func (runeType) IndexOf(s []rune, n rune) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (runeType) LastIndexOf(s []rune, n rune) int {
	if len(s) < 1 {
		return -1
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == n {
			return i
		}
	}
	return -1
}

// Equal 两个slice完全相同, 元素顺序也相同
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

// SameElements 两个slice元素相同，不要求顺序
func (runeType) SameElements(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[rune]int)
	for _, v := range a {
		tmp[v]++
	}

	for _, v := range b {
		tmp[v]--
		if tmp[v] == 0 {
			delete(tmp, v)
		}
	}

	return len(tmp) == 0
}

// Contains 子集, is b subset of a?
func (runeType) Contains(a, b []rune) bool {
	tmp := make(map[rune]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := tmp[v]; !ok {
			return false
		}
	}
	return true
}

// Intersection 交集 a & b，return elements in a and b at same time
func (runeType) Intersection(a, b []rune) []rune {
	tmp := make(map[rune]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	var result []rune
	for _, v := range b {
		if _, ok := tmp[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

// Union 并集 a | b, return elements in a or in b
func (runeType) Union(a, b []rune) []rune {
	tmp := make(map[rune]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		tmp[v] = struct{}{}
	}

	var result []rune
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// Difference 差集 a - diff, return elements in a but not in diff
func (runeType) Difference(a, diff []rune) []rune {
	tmp := make(map[rune]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range diff {
		delete(tmp, v)
	}

	var result []rune
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// SymmetricDifference 对称差集, elements in a or in b, but not in a&b
func (runeType) SymmetricDifference(a, b []rune) []rune {
	tmp := make(map[rune]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := tmp[v]; !ok {
			tmp[v] = struct{}{}
		} else {
			delete(tmp, v)
		}
	}

	var result []rune
	for k := range tmp {
		result = append(result, k)
	}

	return result
}
