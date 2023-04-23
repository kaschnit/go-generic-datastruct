package compare

import "golang.org/x/exp/constraints"

type Priority int

const (
	PriorityRightHigher Priority = -1
	PriorityEqual       Priority = 0
	PriorityLeftHigher  Priority = 1
)

type Comparator[T any] func(left T, right T) Priority

func OrderedComparator[T constraints.Ordered](left T, right T) Priority {
	if left > right {
		return PriorityLeftHigher
	}
	if left < right {
		return PriorityRightHigher
	}
	return PriorityEqual
}
