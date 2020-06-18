package diffsquares

import "math"

func SquareOfSum(n int) int  {
	sum := 0
	for ; n > 0; n-- {
		sum += n
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int  {
	sum := 0
	for ; n > 0; n-- {
		sum += int(math.Pow(float64(n), 2))
	}
	return sum
}

func Difference(n int) int  {
	return SquareOfSum(n) - SumOfSquares(n)
}