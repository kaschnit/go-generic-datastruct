package compare

type Priority int

const (
	PriorityRightHigher Priority = -1
	PriorityEqual       Priority = 0
	PriorityLeftHigher  Priority = 1
)

type Comparator[T any] func(left T, right T) Priority
