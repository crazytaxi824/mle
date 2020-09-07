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
