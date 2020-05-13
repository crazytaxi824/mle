package slice

type byteType struct{}

func Bytes() byteType { return byteType{} }

func (byteType) IndexOf(s []byte, n byte) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (byteType) LastIndexOf(s []byte, n byte) int {
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
func (byteType) Equal(a, b []byte) bool {
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
func (byteType) SameElements(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[byte]int)
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

// Subset 子集, is b subset of a?
func (byteType) Subset(a, b []byte) bool {
	tmp := make(map[byte]struct{})
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
func (byteType) Intersection(a, b []byte) []byte {
	tmp := make(map[byte]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	var result []byte
	for _, v := range b {
		if _, ok := tmp[v]; ok {
			result = append(result, v)
		}
	}

	return result
}

// Union 并集 a | b, return elements in a or in b
func (byteType) Union(a, b []byte) []byte {
	tmp := make(map[byte]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range b {
		tmp[v] = struct{}{}
	}

	var result []byte
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// Difference 差集 a - diff, return elements in a but not in diff
func (byteType) Difference(a, diff []byte) []byte {
	tmp := make(map[byte]struct{})
	for _, v := range a {
		tmp[v] = struct{}{}
	}

	for _, v := range diff {
		delete(tmp, v)
	}

	var result []byte
	for k := range tmp {
		result = append(result, k)
	}

	return result
}

// SymmetricDifference 对称差集, elements in a or in b, but not in a&b
func (byteType) SymmetricDifference(a, b []byte) []byte {
	tmp := make(map[byte]struct{})
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

	var result []byte
	for k := range tmp {
		result = append(result, k)
	}

	return result
}
