// permutation P(nm) = n!/(n-m)!
package pc

import (
	"math/big"
)

func upperPermutation(n, m int64) []*big.Int {
	var upper []*big.Int

	for i := n; i > n-m; i-- {
		upper = append(upper, big.NewInt(i))
	}

	return upper
}

func Permutation(n, m int64) *big.Int {
	u := upperPermutation(n, m)
	return multiAll(u)
}
