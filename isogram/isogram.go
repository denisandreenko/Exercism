package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	for _, s := range input {
		if unicode.IsLetter(s) && strings.Count(input, string(s)) > 1 {
			return false
		}
	}
	return true
}