package grains

import (
	"errors"
)

// Square return the number of grain in each square
func Square(q int) (uint64, error) {
	if q <= 0 || q > 64 {
		return uint64(0), errors.New("Invalid input")
	}

	return 1 << (q - 1), nil
}

// Total return the total number of grains
func Total() uint64 {
	return 1<<64 - 1
}
