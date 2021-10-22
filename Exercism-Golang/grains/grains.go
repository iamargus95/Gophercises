package grains

import "errors"

const maxUint64 = (1 << 64) - 1

func Square(N int) (uint64, error) {
	if N < 1 || N > 64 {
		return uint64(0), errors.New("invalid input")
	}

	return uint64(1 << uint(N-1)), nil
}

func Total() uint64 {
	return uint64(maxUint64)
}
