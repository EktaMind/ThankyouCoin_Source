package utils

import "math/big"

// ToThx number of THX to Wei
func ToThx(thx uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(thx), big.NewInt(1e18))
}
