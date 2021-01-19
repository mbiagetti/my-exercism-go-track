package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, j := range word {
		if unicode.IsLetter(j) && strings.Count(word, string(word[i])) > 1 {
			return false
		}
	}
	return true
}
