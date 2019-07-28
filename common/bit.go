package common

import "math/bits"

func FirstOne(src uint64) uint {
	return uint(bits.TrailingZeros64(src))
}
