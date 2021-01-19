package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	tot := len(rhyme)
	out := make([]string, tot)

	for idx := range rhyme {
		if idx+1 == tot { // last element
			out[idx] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			out[idx] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[idx], rhyme[idx+1])
		}
	}

	return out
}
