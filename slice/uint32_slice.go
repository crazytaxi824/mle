package slice

type uint32Type struct{}

func Uint32() uint32Type { return uint32Type{} }

func (uint32Type) IndexOf(s []uint32, n uint32) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (uint32Type) LastIndexOf(s []uint32, n uint32) int {
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

// SameElements 两个slice元素相同，不要求顺序
func (uint32Type) SameElements(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[uint32]int)
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
func (uint32Type) Contains(a, b []uint32) bool {
	tmp := make(map[uint32]struct{})
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
func (uint32Type) Intersection(a, b []uint32) []uint32 {
	tmp := make(map[uint32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	var result []uint32
	for _, v := range b {
		if _, ok := tmp[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

// Union 并集 a | b, return elements in a or in b
func (uint32Type) Union(a, b []uint32) []uint32 {
	tmp := make(map[uint32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		tmp[v] = struct{}{}
	}

	var result []uint32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// Difference 差集 a - diff, return elements in a but not in diff
func (uint32Type) Difference(a, diff []uint32) []uint32 {
	tmp := make(map[uint32]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range diff {
		delete(tmp, v)
	}

	var result []uint32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// SymmetricDifference 对称差集, elements in a or in b, but not in a&b
func (uint32Type) SymmetricDifference(a, b []uint32) []uint32 {
	tmp := make(map[uint32]struct{})
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

	var result []uint32
	for k := range tmp {
		result = append(result, k)
	}

	return result
}
