package compare

import "golang.org/x/exp/constraints"

// Priority is used to indicate a result of a comparison by a comparator.
type Priority int

const (
	// PriorityRightHigher indicates the right hand side value is higher priority.
	PriorityRightHigher Priority = -1
	// PriorityEqual indicates the left and right hand side values are equal in priority.
	PriorityEqual Priority = 0
	// PriorityLeftHigher indicates the left hand side value is higher priority.
	PriorityLeftHigher Priority = 1
)

// Comparator is a function that defines the ordering of any two types, which may not be orderable
// by default. It allows defining arbitrary comparisons and ordering for data structure that require
// comparison and ordering.
type Comparator[T any] func(left T, right T) Priority

// OrderedComparator is a comparator that provides standard ordering (i.e., =, <, and > operators)
// to any orderable type.
func OrderedComparator[T constraints.Ordered](left T, right T) Priority {
	if left > right {
		return PriorityLeftHigher
	}
	if left < right {
		return PriorityRightHigher
	}
	return PriorityEqual
}

// OrderedComparator is a comparator that provides the opposite of standard ordering
// (i.e., =, <, and > operators) to any orderable type.
func OppositeOrderedComparator[T constraints.Ordered](left T, right T) Priority {
	if left < right {
		return PriorityLeftHigher
	}
	if left > right {
		return PriorityRightHigher
	}
	return PriorityEqual
}

// Opposite can be used to transform a comparator into a comparator that returns the "opposite"
// result, meaning left and high priority is swapped.
func Opposite[T any](c Comparator[T]) Comparator[T] {
	return func(left, right T) Priority {
		cmp := c(left, right)
		switch cmp {
		case PriorityLeftHigher:
			return PriorityRightHigher
		case PriorityRightHigher:
			return PriorityLeftHigher
		default:
			return PriorityEqual
		}
	}
}
