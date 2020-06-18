package pangram

import "strings"

var alphabet = []rune("abcdefghijklmnopqrstuvwxyz")

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for _, val := range alphabet {
		if !strings.ContainsRune(input, val) {
			return false
		}
	}
	return true
}