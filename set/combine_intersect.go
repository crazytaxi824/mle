package set

// 交集
type intersection struct{}

func Intersect() intersection {
	return intersection{}
}

func (intersection) Int16(a, b *int16HashSet) *int16HashSet {
	var r = NewInt16Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

