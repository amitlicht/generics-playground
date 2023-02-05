package filters

// FilterStrings filters a slice of strings
func FilterStrings(items []string, pred func(item string, i int) bool) []string {
	r := make([]string, 0)
	for i, item := range items {
		if pred(item, i) {
			r = append(r, item)
		}
	}
	return r
}

// FilterInts filters a slice of integers
func FilterInts(items []int, pred func(item int, i int) bool) []int {
	r := make([]int, 0)
	for i, item := range items {
		if pred(item, i) {
			r = append(r, item)
		}
	}
	return r
}
