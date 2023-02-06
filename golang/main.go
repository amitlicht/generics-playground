package main

import (
	"errors"
	"fmt"
	"github.com/amitlicht/generics-playground/filters"
	"github.com/amitlicht/generics-playground/set"
	"github.com/samber/lo"
	"sort"
	"strings"
)

var (
	animals = []string{"gopher", "otter", "mole", "snake"}
	isCool  = func(item string, i int) bool { return strings.HasSuffix(item, "er") }

	numbers = []int{0, 100, 50, -10, -365}
	isPos   = func(item int, i int) bool { return item > 0 }
)

// shuffle returns a random shuffling of its input string
func shuffle(a string) string {
	return string(lo.Shuffle([]rune(a)))
}

// log prints value to screen
func log(v any) {
	fmt.Printf("%v\n", v)
}

// withTypedFilters demonstrates the use of typed filters
func withTypedFilters() {
	log("--- With typed filters: ---")
	log(filters.FilterStrings(animals, isCool)) // [gopher otter]
	log(filters.FilterInts(numbers, isPos))     // [100 50]
}

// withAnyFilter demonstrates the use of "any" filters
func withAnyFilter() {
	log("--- With FilterAny: ---")

	// convert animals from []string to []any
	animalsAny := make([]any, len(animals))
	for i, a := range animals {
		animalsAny[i] = a
	}
	isCoolAny := func(item any, i int) bool { return isCool(item.(string), i) }
	log(filters.FilterAny(animalsAny, isCoolAny)) // [gopher otter]

	// convert numbers from []int to []any
	numbersAny := make([]any, len(numbers))
	for i, n := range numbers {
		numbersAny[i] = n
	}
	isPosAny := func(item any, i int) bool { return isPos(item.(int), i) }
	log(filters.FilterAny(numbersAny, isPosAny)) // [100 50]

	//filters.FilterAny(animalsAny, isPosAny) // panic: interface conversion: interface {} is string, not int
}

// withReflectionFilter demonstrates the use of reflection-based filters
func withReflectionFilter() {
	log("--- With FilterReflection: ---")
	log(filters.FilterReflection(animals, isCool)) // [gopher otter]
	log(filters.FilterReflection(numbers, isPos))  // [100 50]

	//filters.FilterReflection(animals, isPos) // panic: reflect: Call using string as type int
}

// withGenericFilter demonstrates the use of generic filters
func withGenericFilter() {
	log("--- With FilterGeneric: ---")
	log(filters.FilterGeneric[string](animals, isCool)) // [gopher otter]
	log(filters.FilterGeneric[int](numbers, isPos))     // [100 50]

	// Type inference
	log(filters.FilterGeneric(animals, isCool)) // [gopher otter]
	log(filters.FilterGeneric(numbers, isPos))  // [100 50]

	// Compiler error:
	// Cannot use 'isPos' (type func(item int, i int) bool) as the type func(item string, i int) bool
	// filters.FilterGeneric(animals, isPos)
}

// withLoFilter demonstrates the use of lo.Filter
func withLoFilter() {
	log("--- With lo.Filter: ---")
	log(lo.Filter(animals, isCool)) // [gopher otter]
	log(lo.Filter(numbers, isPos))  // [100 50]
}

// withNestedLo demonstrates the use of lo functionality
func withNestedLo() {
	log("--- With nested lo: ---")
	// map cool animals to random shuffling of their names
	log(
		lo.FromPairs(
			lo.Map(
				lo.Filter(animals, isCool),
				func(a string, i int) lo.Entry[string, string] {
					return lo.Entry[string, string]{
						Key:   a,
						Value: shuffle(a),
					}
				},
			),
		),
	)
}

// withLoFilterErr demonstrates the use of FilterErr
func withLoFilterErr() {
	log("--- With FilterErr: ---")
	animals := []string{"gopher", "otter", "mole", "snake", ""}
	isCoolErr := func(item string, i int) (bool, error) {
		if item == "" {
			return false, errors.New("anonymous animal")
		}
		return isCool(item, i), nil
	}

	if _, err := filters.FilterErr(animals, isCoolErr); err != nil {
		log(err)
	}
}

// withSets demonstrates the use of the set package
func withSets() {
	log("--- With Set: ---")
	s := set.NewSet[string]()
	s.Add("gopher", "otter", "mole", "snake")
	if s.Contains("snake") {
		s.Remove("mole") // they're scared of snakes
	}

	s.ForEach(func(a string) {
		// random shuffle the animal's name
		log(shuffle(a))
	})
}

// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.s) }
func (s SliceFn[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }
func (s SliceFn[T]) Less(i, j int) bool { return s.less(s.s[i], s.s[j]) }

// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, less})
}

func main() {
	// different filter implementations
	withTypedFilters()
	withAnyFilter()
	withReflectionFilter()
	withGenericFilter()
	withLoFilter()

	// list comprehension in go
	withNestedLo()

	// filter + err
	withLoFilterErr()

	// set usage example
	withSets()
}
