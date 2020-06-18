package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(str []string) FreqMap {
	channel := make(chan FreqMap)
	result := FreqMap{}
	for _, s := range str {
		go func(s string) {
			channel <- Frequency(s)
		}(s)
	}
	for i := 0; i < len(str); i++ {
		goRes := <- channel
		for k, v := range goRes {
			result[k] += v
		}
	}
	return result
}
