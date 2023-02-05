package set

/**
basic generic Set implementation, inspired by github.com/amit7itz/goset
*/

// Set is a simple set implementation using map as its store
type Set[T comparable] map[T]any

// NewSet returns a new Set of the given items
func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Add adds item(s) to the Set
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

// Remove removes item(s) from the Set if they exist
func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

// Contains returns whether an item is in the Set
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// ForEach runs a function on all the items in the Set
func (s Set[T]) ForEach(f func(item T)) {
	for item := range s {
		f(item)
	}
}
