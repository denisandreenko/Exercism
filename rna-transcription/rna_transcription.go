package strand

var rnaMap = map[rune]rune {
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	result := make([]rune, 0)
	for _, v := range dna {
		result = append(result, rnaMap[v])
	}
	return string(result)
}
