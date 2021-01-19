package acronym

import (
	"strings"
	u "unicode"
)

func Abbreviate(s string) string {
	f := func(c rune) bool {
		return u.IsSpace(c) || c == '-'
	}
	var arr = strings.FieldsFunc(s, f)
	s = ""
	for _, k := range arr {
		s += string([]rune(k)[0])
	}
	return strings.ToUpper(s)
}

func sortString(s string) string {
	result := ""

	a := "A"
	for _, char := range s {
		if char < a {
			a = char
		}
	}

	result = result + a

	b := "A"
	for _, char := range s {
		if char < b || char < a {
			b = char
		}

		result = result + b

	}