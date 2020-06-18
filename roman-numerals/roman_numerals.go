package romannumerals

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var romanNumerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(number int) (string, error) {
	var romanNumeral string

	if number < 1  || number > 3000 {
		return "", fmt.Errorf("error")
	}

	numberString := strconv.Itoa(number)
	digits := len(numberString)

	for i, digitString := range numberString {
		place := int(math.Pow(10, float64(digits-i-1)))
		digit, err := strconv.Atoi(string(digitString))
		if err != nil {
			return "", err
		}
		if digit < 4 {
			romanNumeral += strings.Repeat(romanNumerals[place], digit)
		} else if digit == 4 {
			romanNumeral += romanNumerals[place] + romanNumerals[place*5]
		} else if digit > 4 && digit < 9 {
			romanNumeral += romanNumerals[place*5] + strings.Repeat(romanNumerals[place], digit%5)
		} else if digit == 9 {
			romanNumeral += romanNumerals[place] + romanNumerals[int(math.Pow(10, float64(digits-i)))]
		}
	}
	return romanNumeral, nil
}