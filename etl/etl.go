package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for k, v := range input {
		for _, s := range v {
			result[strings.ToLower(s)] = k
		}
	}
	return result
}