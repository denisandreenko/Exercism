package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	input = strings.ToLower(input)
	wordMap := make(map[string]int)
	regex := regexp.MustCompile("[a-zA-Z0-9']+")
	words := regex.FindAllString(input, -1)
	for _, word := range words {
		word = strings.Trim(word, "'")
		if strings.Contains(input, word) {
			wordMap[word] += 1
		}
	}
	return wordMap
}