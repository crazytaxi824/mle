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

func (union) Int32(a, b *int32HashSet) *int32HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Int64(a, b *int64HashSet) *int64HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Int(a, b *intHashSet) *intHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Uint16(a, b *uint16HashSet) *uint16HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Uint32(a, b *uint32HashSet) *uint32HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Uint64(a, b *uint64HashSet) *uint64HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) String(a, b *stringHashSet) *stringHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Rune(a, b *runeHashSet) *runeHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}

func (union) Interface(a, b *interfaceHashSet) *interfaceHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Add(k)
	}
	return r
}
