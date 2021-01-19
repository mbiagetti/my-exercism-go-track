package reverse

// String is used to reverse a string
func String(s string) string {
	var result []rune
	for _, r := range s {
		result = append([]rune{r}, result...)
	}
	return string(result)
}
