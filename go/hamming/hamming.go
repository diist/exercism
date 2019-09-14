package hamming

import "errors"

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strands lengths cannot be different")
	}

	var d int
	for i, c := range a {
		if c != rune(b[i]) {
			d++
		}
	}

	return d, nil
}
