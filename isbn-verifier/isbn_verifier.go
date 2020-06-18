package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(input string) bool {
	input = strings.ReplaceAll(input, "-", "")
	if len(input) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 10; i++ {
		if !unicode.IsNumber(rune(input[i])) && !(i == 9 && input[i] == 'X') {
			return false
		}
		if input[i] == 'X' {
			sum += 10
		} else {
			sum += int(input[i] - '0') * (10 - i)
		}
	}
	if sum % 11 == 0 {
		return true
	}
	return false
}