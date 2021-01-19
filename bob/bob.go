package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.Trim(remark, "\t")
	remark = strings.Trim(remark, "\n")
	remark = strings.Trim(remark, "\r")
	remark = strings.TrimSpace(remark)

	var isUppercase = strings.Compare(remark, strings.ToUpper(remark)) == 0
	var isQuestion = strings.HasSuffix(remark, "?")

	switch {
	case isUppercase && isQuestion && haveLetter(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isUppercase && haveLetter(remark):
		return "Whoa, chill out!"
	case strings.Compare(remark, "?") == 0 || len(remark) == 0:
		return "Fine. Be that way!"

	default:
		return "Whatever."
	}
}

func haveLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
