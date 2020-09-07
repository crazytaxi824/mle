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

func (intersection) Int32(a, b *int32HashSet) *int32HashSet {
	var r = NewInt32Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Int64(a, b *int64HashSet) *int64HashSet {
	var r = NewInt64Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Int(a, b *intHashSet) *intHashSet {
	var r = NewIntSet()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Uint16(a, b *uint16HashSet) *uint16HashSet {
	var r = NewUint16Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Uint32(a, b *uint32HashSet) *uint32HashSet {
	var r = NewUint32Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Uint64(a, b *uint64HashSet) *uint64HashSet {
	var r = NewUint64Set()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Rune(a, b *runeHashSet) *runeHashSet {
	var r = NewRuneSet()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) Interface(a, b *interfaceHashSet) *interfaceHashSet {
	var r = NewInterSet()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}

func (intersection) String(a, b *stringHashSet) *stringHashSet {
	var r = NewStringSet()
	for k := range b.elements {
		if a.Contains(k) {
			r.Add(k)
		}
	}
	return r
}
