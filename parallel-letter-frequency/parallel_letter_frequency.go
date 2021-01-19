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

// ConcurrentFrequency count frequency in a parallel way
func ConcurrentFrequency(strings []string) interface{} {
	f := FreqMap{}
	messages := make(chan FreqMap, 10)

	for _, j := range strings {
		go func(i string) {
			messages <- Frequency(i)
		}(j)
	}

	for range strings {
		for k, v := range <-messages {
			f[k] += v
		}
	}

	return f
}
