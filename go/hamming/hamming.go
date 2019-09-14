package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var d int

	if len(a) != len(b) {
		return -1, errors.New("Strands lengths cannot be different")
	}

	for i, c := range a {
		if c != rune(b[i]) {
			d++
		}
	}

	return d, nil
}
