package set

// ISet is the generic interface for types implementing set operations
type ISet[T comparable] interface {
	Add(items ...T)
	Remove(items ...T)
	Contains(item T) bool
	ForEach(f func(item T))
}
