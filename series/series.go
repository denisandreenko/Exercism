package series

func All(n int, s string) []string {
	res := make([]string, 0)
	for i := 0; i < len(s) - n + 1; i++ {
		res = append(res, UnsafeFirst(n, s[i:]))
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return s[:len(s)]
	}
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) >= n {
		return s[:n], true
	}
	return
}