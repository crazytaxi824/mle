// combination C(nm) = n!/m!(n-m)!
package pc

import (
	"math/big"
)

func maxAndMin(a, b int64) (int64, int64) {
	if a > b {
		return a, b
	}
	return b, a
}

func upperAndLowerCombination(n, m int64) ([]*big.Int, []*big.Int) {
	var upper, lower []*big.Int
	maxNum, minNum := maxAndMin(m, n-m)

	for i := n; i > maxNum; i-- {
		upper = append(upper, big.NewInt(i))
	}

	for i := int64(1); i <= minNum; i++ {
		lower = append(lower, big.NewInt(i))
	}

	return upper, lower
}

func multiAll(a []*big.Int) *big.Int {
	r := big.NewInt(1)
	for i := range a {
		r.Mul(r, a[i])
	}
	return r
}

func Combination(n, m int64) *big.Int {
	u, l := upperAndLowerCombination(n, m)
	up := multiAll(u)
	low := multiAll(l)

	r := new(big.Int)
	r.Div(up, low)

	return r
}
