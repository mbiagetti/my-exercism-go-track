package twofer

func ShareWith(name string) string {
	me := "you"
	if len(name) > 0 {
		me = name
	}

	return "One for " + me + ", one for me."

}
