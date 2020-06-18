package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	runes := []rune(input)
	sum := 0
	if len(input) <= 1 {
		return false
	}
	for i, elem := range input {
		if !unicode.IsNumber(elem) {
			return false
		}
		intVal := int(runes[i] - '0')
		if (len(input) % 2 == 1 && i % 2 == 1) || (len(input) % 2 == 0 && i % 2 == 0)  {
			intVal *= 2
			if intVal > 9 {
				intVal -= 9
			}
		}
		sum += intVal
	}
	if sum % 10 == 0 {
		return true
	}
	return false
}