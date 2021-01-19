package collatzconjecture

import 	"errors"


func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("input must be greater than zero")
	}
	step := 0

	for ; input != 1; step++ {
		if input%2 == 0 {
			input /= 2
		} else {
			input = input*3 + 1
		}
	}

	return step, nil
}
