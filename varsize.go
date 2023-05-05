package varsize

import (
	"math/bits"
)

// Int computes a number of bytes required to write off a varint
// encoding of the given signed integer value x.
func Int[T Signed](x T) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}

	return ((bits.Len64(ux | 1)) + 6) / 7
}

// Uint computes a number of bytes required to write off an uvarint
// encoding of the given unsigned integer value x.
func Uint[T Unsigned](x T) int {
	return ((bits.Len64(uint64(x) | 1)) + 6) / 7
}

// Len is just a shortcut for `Uint[uint64]` for lengths of slices converted to uint64.
func Len[T any](v []T) int {
	return (bits.Len64(uint64(len(v))|1) + 6) / 7
}
