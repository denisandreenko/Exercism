package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	answer := ""
	remark = strings.TrimSpace(remark)
	if IsUpper(remark) && IsQuestion(remark) {
		answer += "Calm down, I know what I'm doing!"
	} else if IsUpper(remark) {
		answer += "Whoa, chill out!"
	} else if IsQuestion(remark) {
		answer += "Sure."
	} else if remark == "" {
		answer += "Fine. Be that way!"
	} else {
		answer += "Whatever."
	}
	return answer
}

func IsUpper(s string) bool {
	if !HasLetters(s) {
		return false
	}
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func HasLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func IsQuestion(s string) bool {
	if len(s) == 0 {
		return false
	}
	return string(s[len(s) - 1]) == "?"
}
