// Package letter is the package for the Exercism exercise letter
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

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap in a concurrent fashion using channels
func ConcurrentFrequency(ss []string) FreqMap {
	var fm = FreqMap{}
	ch := make(chan FreqMap, 10)

	for _, s := range ss {
		go func(s string, ch chan FreqMap) {
			ch <- Frequency(s)
		}(s, ch)
	}

	for range ss {
		for k, v := range <-ch {
			fm[k] += v
		}
	}

	return fm
}
