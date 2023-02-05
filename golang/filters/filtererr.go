package filters

// FilterErr is similar to Filter, with error handling for the pred function
// See https://github.com/samber/lo/pull/292
func FilterErr[T any](items []T, pred func(item T, index int) (bool, error)) ([]T, error) {
	var result []T

	for i, item := range items {
		if res, err := pred(item, i); err != nil {
			return nil, err
		} else if res {
			result = append(result, item)
		}
	}

	return result, nil
}
