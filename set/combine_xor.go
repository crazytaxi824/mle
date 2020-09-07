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
