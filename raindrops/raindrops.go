package raindrops

import (
	"strconv"
	"strings"
)

func Convert(input int) string  {
	result := make([]string, 0)
	if input % 3 == 0 {
		result = append(result, "Pling")
	}
	if input % 5 == 0 {
		result = append(result, "Plang")
	}
	if input % 7 == 0 {
		result = append(result, "Plong")
	}
	if len(result) == 0 {
		result = append(result, strconv.Itoa(input))
	}
	return strings.Join(result, "")
}