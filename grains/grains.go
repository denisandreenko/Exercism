package grains

import (
	"errors"
	"math"
)

const maxSquares  = 64

func Square(input int) (uint64, error) {
	if input <= 0 || input > maxSquares {
		return 0, errors.New("error")
	}
	return uint64(math.Pow(2, float64(input - 1))), nil
}

func Total() uint64 {
	var totalScore uint64 = 0
	for i := 1; i <= maxSquares; i++ {
		res, err := Square(i)
		if err == nil {
			totalScore += res
		}
	}
	return totalScore
}