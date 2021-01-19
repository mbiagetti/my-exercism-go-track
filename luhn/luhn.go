package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(s string) bool {
	// remove spaces
	s = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)

	if len(s) < 2 {
		return false
	}

	sum := 0
	for i, j := range s {
		x, _ := strconv.Atoi(string(j))
		if (i+1)%2 == 0 {
			a := x * 2
			if a > 9 {
				a -= 9
			}
			sum += a
		} else {
			sum += x
		}
	}
	return sum%10 == 0
}
