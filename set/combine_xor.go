package set

// 对称差集
type xor struct{}

func XOR() xor {
	return xor{}
}

func (xor) Int16(a, b *int16HashSet) *int16HashSet {
	u := Union().Int16(a, b)
	i := Intersect().Int16(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Int32(a, b *int32HashSet) *int32HashSet {
	u := Union().Int32(a, b)
	i := Intersect().Int32(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Int64(a, b *int64HashSet) *int64HashSet {
	u := Union().Int64(a, b)
	i := Intersect().Int64(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Int(a, b *intHashSet) *intHashSet {
	u := Union().Int(a, b)
	i := Intersect().Int(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Uint16(a, b *uint16HashSet) *uint16HashSet {
	u := Union().Uint16(a, b)
	i := Intersect().Uint16(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Uint32(a, b *uint32HashSet) *uint32HashSet {
	u := Union().Uint32(a, b)
	i := Intersect().Uint32(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Uint64(a, b *uint64HashSet) *uint64HashSet {
	u := Union().Uint64(a, b)
	i := Intersect().Uint64(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) String(a, b *stringHashSet) *stringHashSet {
	u := Union().String(a, b)
	i := Intersect().String(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Rune(a, b *runeHashSet) *runeHashSet {
	u := Union().Rune(a, b)
	i := Intersect().Rune(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}

func (xor) Interface(a, b *interfaceHashSet) *interfaceHashSet {
	u := Union().Interface(a, b)
	i := Intersect().Interface(a, b)

	for k := range i.elements {
		u.Delete(k)
	}
	return u
}
