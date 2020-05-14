package slice

type int16Type struct{}

func Int16() int16Type { return int16Type{} }

func (int16Type) IndexOf(s []int16, n int16) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (int16Type) LastIndexOf(s []int16, n int16) int {
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
func (int16Type) Equal(a, b []int16) bool {
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

// SameElements 两个slice元素相同，不要求顺序，但是重复次数必须相同
func (int16Type) SameElements(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}

	tmp := make(map[int16]int, len(a))
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

// Contains a包含b中的所有元素，不考虑顺序和重复次数
func (int16Type) Contains(a, b []int16) bool {
	tmp := make(map[int16]struct{}, len(a))
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

// // Contains any of the elements
// func (int16Type) ContainsAny(a, b []int16) bool {
// 	tmp := make(map[int16]struct{})
// 	for _, v := range a {
// 		tmp[v] = struct{}{}
// 	}
//
// 	for _, v := range b {
// 		if _, ok := tmp[v]; ok {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// // Intersection 交集 a & b，return elements in a and b at same time
// func (int16Type) Intersection(a, b []int16) []int16 {
// 	tmp := make(map[int16]struct{})
// 	for _, v := range a {
// 		tmp[v] = struct{}{}
// 	}
//
// 	var result []int16
// 	for _, v := range b {
// 		if _, ok := tmp[v]; ok {
// 			result = append(result, v)
// 		}
// 	}
//
// 	return result
// }
//
// // Union 并集 a | b, return elements in a or in b
// func (int16Type) Union(a, b []int16) []int16 {
// 	tmp := make(map[int16]struct{})
// 	for _, v := range a {
// 		tmp[v] = struct{}{}
// 	}
//
// 	for _, v := range b {
// 		tmp[v] = struct{}{}
// 	}
//
// 	var result []int16
// 	for k := range tmp {
// 		result = append(result, k)
// 	}
//
// 	return result
// }
//
// // Difference 差集 a - diff, return elements in a but not in diff
// func (int16Type) Difference(a, diff []int16) []int16 {
// 	tmp := make(map[int16]struct{})
// 	for _, v := range a {
// 		tmp[v] = struct{}{}
// 	}
//
// 	for _, v := range diff {
// 		delete(tmp, v)
// 	}
//
// 	var result []int16
// 	for k := range tmp {
// 		result = append(result, k)
// 	}
//
// 	return result
// }
//
// // SymmetricDifference 对称差集, elements in a or in b, but not in a&b
// func (int16Type) SymmetricDifference(a, b []int16) []int16 {
// 	tmp := make(map[int16]struct{})
// 	for _, v := range a {
// 		tmp[v] = struct{}{}
// 	}
//
// 	for _, v := range b {
// 		if _, ok := tmp[v]; !ok {
// 			tmp[v] = struct{}{}
// 		} else {
// 			delete(tmp, v)
// 		}
// 	}
//
// 	var result []int16
// 	for k := range tmp {
// 		result = append(result, k)
// 	}
//
// 	return result
// }
