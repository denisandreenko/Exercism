package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(ibase int, digits []int, obase int) ([]int, error) {
	if ibase < 2 {
		return nil, errors.New("input base must be >= 2")
	} else if obase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	num := 0
	for i:=len(digits)-1; i>=0; i-- {
		if digits[i] < 0 || digits[i] >= ibase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num += digits[i] * int(math.Pow(float64(ibase), float64(len(digits)-1-i)))
	}

	var resDigits []int
	for num > 0 {
		resDigits = append([]int{num%obase}, resDigits...)
		num /= obase
	}
	if len(resDigits) == 0 {
		resDigits = []int{0}
	}

	return resDigits, nil

}