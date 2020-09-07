package set

// 差集
type subtraction struct{}

func Subtract() subtraction {
	return subtraction{}
}

// Int16 is subtraction (a - b)
// subtraction a - b and b - a is not the same.
func (subtraction) Int16(a, b *int16HashSet) *int16HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Int32(a, b *int32HashSet) *int32HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Int64(a, b *int64HashSet) *int64HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Int(a, b *intHashSet) *intHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Uint16(a, b *uint16HashSet) *uint16HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Uint32(a, b *uint32HashSet) *uint32HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Uint64(a, b *uint64HashSet) *uint64HashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) String(a, b *stringHashSet) *stringHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Interface(a, b *interfaceHashSet) *interfaceHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}

func (subtraction) Rune(a, b *runeHashSet) *runeHashSet {
	r := a.Copy()
	for k := range b.elements {
		r.Delete(k)
	}
	return r
}
