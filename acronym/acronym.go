package acronym

import (
	"regexp"
	"unicode"
)

func Abbreviate(s string) string {
	regex := regexp.MustCompile("[ !,.\\-_=]")
	words := regex.Split(s, -1)
	answer := make([]rune, 0)
	for _, word := range words {
		for _, sym := range word {
			if unicode.IsLetter(sym) {
				answer = append(answer, unicode.ToUpper(sym))
				break
			}
		}
	}
	return string(answer)
}
