package flatten

func Flatten(input interface{}) []interface{} {
	return flat(input)
}

func flat(data interface{}) []interface{} {
	out := []interface{}{}
	switch data.(type) {
	case []interface{}:
		for _, i := range data.([]interface{}) {
			if i != nil {
				out = append(out, flat(i)...)
			}
		}
	default:
		out = append(out, data)
	}

	return out
}
