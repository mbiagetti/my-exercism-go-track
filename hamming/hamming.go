package hamming

import "errors"

func Distance(a, b string) (int, error) {

	i := 0
	if len(a) != len(b) {
		return -1, errors.New("different length")
	}

	for j := 0; j < len(a); j++ {
		if a[j] != b[j] {
			i += 1
		}
	}

	return i, nil
}
