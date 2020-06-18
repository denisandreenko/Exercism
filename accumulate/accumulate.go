package accumulate

func Accumulate(input []string, f func(string) string) []string {
	for i := range input {
		input[i] = f(input[i])
	}
	return input
}