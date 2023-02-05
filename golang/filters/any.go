package filters

// FilterAny filters a slice of any (aka interface{})
func FilterAny(items []any, pred func(item any, i int) bool) []any {
	r := make([]any, 0)
	for i, item := range items {
		if pred(item, i) {
			r = append(r, item)
		}
	}
	return r
}
