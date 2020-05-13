package slice

type int32Type struct{}

func Int32() int32Type { return int32Type{} }

func (int32Type) IndexOf(s []int32, n int32) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (int32Type) LastIndexOf(s []int32, n int32) int {
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
func (int32Type) Equal(a, b []int32) bool {
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
func (int32Type) SameElements(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[int32]int)
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
func (int32Type) Contains(a, b []int32) bool {
	tmp := make(map[int32]struct{})
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
func (int32Type) ContainsAny(a, b []int32) bool {
	tmp := make(map[int32]struct{})
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
func (int32Type) Intersection(a, b []int32) []int32 {
	tmp := make(map[int32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	var result []int32
	for _, v := range b {
		if _, ok := tmp[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

// Union 并集 a | b, return elements in a or in b
func (int32Type) Union(a, b []int32) []int32 {
	tmp := make(map[int32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		tmp[v] = struct{}{}
	}

	var result []int32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// Difference 差集 a - diff, return elements in a but not in diff
func (int32Type) Difference(a, diff []int32) []int32 {
	tmp := make(map[int32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range diff {
		delete(tmp, v)
	}

	var result []int32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// SymmetricDifference 对称差集, elements in a or in b, but not in a&b
func (int32Type) SymmetricDifference(a, b []int32) []int32 {
	tmp := make(map[int32]struct{})
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

	var result []int32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}
