package slice

type intType struct{}

func Int() intType { return intType{} }

func (intType) IndexOf(s []int, n int) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (intType) LastIndexOf(s []int, n int) int {
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

// SameElements 两个slice元素相同，不要求顺序
func (intType) SameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[int]int)
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
func (intType) Contains(a, b []int) bool {
	tmp := make(map[int]struct{})
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

// Contains any of the elements
func (intType) ContainsAny(a, b []int) bool {
	tmp := make(map[int]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := tmp[v]; ok {
			return true
		}
	}
	return false
}

// Intersection 交集 a & b，return elements in a and b at same time
func (intType) Intersection(a, b []int) []int {
	tmp := make(map[int]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	var result []int
	for _, v := range b {
		if _, ok := tmp[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

// Union 并集 a | b, return elements in a or in b
func (intType) Union(a, b []int) []int {
	tmp := make(map[int]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		tmp[v] = struct{}{}
	}

	var result []int
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// Difference 差集 a - diff, return elements in a but not in diff
func (intType) Difference(a, diff []int) []int {
	tmp := make(map[int]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range diff {
		delete(tmp, v)
	}

	var result []int
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// SymmetricDifference 对称差集, elements in a or in b, but not in a&b
func (intType) SymmetricDifference(a, b []int) []int {
	tmp := make(map[int]struct{})
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

	var result []int
	for k := range tmp {
		result = append(result, k)
	}

	return result
}
