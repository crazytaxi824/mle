package set

// 并集
type union struct{}

func Union() union {
	return union{}
}

func (union) Int16(a, b *int16HashSet) *int16HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}
