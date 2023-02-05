package filters

// FilterGeneric filters a slice of items of type T
func FilterGeneric[T any](items []T, pred func(item T, i int) bool) []T {
	r := make([]T, 0)
	for i, item := range items {
		if pred(item, i) {
			r = append(r, item)
		}
	}
	return r
}
