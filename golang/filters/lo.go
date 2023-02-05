package filters

import (
	"github.com/samber/lo"
	"strings"
)

func withLoFilter() {
	animals := []string{"gopher", "otter", "mole", "snake"}
	isCool := func(item string, _ int) bool { return strings.HasSuffix(item, "er") }
	lo.Filter(animals, isCool) // [gopher otter]

	numbers := []int{0, 100, 50, -10, -365}
	isPos := func(item int, _ int) bool { return item > 0 }
	lo.Filter(numbers, isPos) // [100 50]
}
