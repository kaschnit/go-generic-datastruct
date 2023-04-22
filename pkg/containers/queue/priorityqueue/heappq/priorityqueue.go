package heappq

import (
	"fmt"
	"strings"
)

type Priority int

const (
	PriorityRightHigher Priority = -1
	PriorityEqual       Priority = 0
	PriorityLeftHigher  Priority = 1
)

type PriorityComparator[T any] func(left T, right T) Priority

type HeapPQ[T any] struct {
	comparator PriorityComparator[T]
	items      []T
}

func New[T any](comparator PriorityComparator[T], values ...T) *HeapPQ[T] {
	q := HeapPQ[T]{
		comparator: comparator,
		items:      make([]T, 1),
	}
	q.PushAll(values...)
	return &q
}

func (q *HeapPQ[T]) Empty() bool {
	return q.Size() == 0
}

func (q *HeapPQ[T]) Size() int {
	return len(q.items) - 1
}

func (q *HeapPQ[T]) Clear() {
	q.items = make([]T, 1)
}

func (q *HeapPQ[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("HeapPQ\n")
	strs := make([]string, 0, q.Size())
	for i := 1; i < len(q.items); i++ {
		strs = append(strs, fmt.Sprintf("%v", q.items[i]))
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (q *HeapPQ[T]) Push(value T) {
	// Push onto the end
	q.items = append(q.items, value)

	// Percolate up to maintain heap invariant
	for fixIndex := len(q.items) - 1; fixIndex != 1; {
		parentIndex := parent(fixIndex)

		// Check whether the parent is higher priority
		isParentInPlace := q.comparator(q.items[fixIndex], q.items[parentIndex]) != PriorityLeftHigher

		// If the parent priority is higher, we're done percolating up.
		if isParentInPlace {
			break
		}

		// Swap the current index with the parent and continue
		q.items[fixIndex], q.items[parentIndex] = q.items[parentIndex], q.items[fixIndex]
		fixIndex = parentIndex
	}
}

func (q *HeapPQ[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *HeapPQ[T]) Pop() (value T, ok bool) {
	value, ok = q.Peek()
	if ok {
		return value, ok
	}

	// Move the last item to the root
	q.items[1], q.items[len(q.items)-1] = q.items[len(q.items)-1], q.items[1]

	// Remove the last item, which was the root
	q.items = q.items[:len(q.items)-1]

	// Percolate down to maintain the heap invariant
	for fixIndex := 1; fixIndex != len(q.items)-1; {
		// Check whether the left child is higher priority
		leftIndex := leftChild(fixIndex)
		if leftIndex < len(q.items) {
			// Child is in place if it's not higher priority than the parent.
			isChildInPlace := q.comparator(q.items[leftIndex], q.items[fixIndex]) != PriorityLeftHigher

			// Swap with the child if the child is not in place.
			if !isChildInPlace {
				q.items[fixIndex], q.items[leftIndex] = q.items[leftIndex], q.items[fixIndex]
				fixIndex = leftIndex
				continue
			}
		}

		// Check whether the right child is higher priority
		rightIndex := rightChild(fixIndex)
		if rightIndex < len(q.items) {
			// Child is in place if it's not higher priority than the parent.
			isChildInPlace := q.comparator(q.items[rightIndex], q.items[fixIndex]) != PriorityLeftHigher

			// Swap with the child if the child is not in place.
			if !isChildInPlace {
				q.items[fixIndex], q.items[rightIndex] = q.items[rightIndex], q.items[fixIndex]
				fixIndex = rightIndex
				continue
			}
		}

		// Node is in place, done.
		break
	}

	return value, ok
}

func (q *HeapPQ[T]) Peek() (value T, ok bool) {
	if q.Empty() {
		return *new(T), false
	}
	return q.items[1], true
}

func parent(index int) int {
	return index / 2
}

func leftChild(index int) int {
	return index * 2
}

func rightChild(index int) int {
	return index*2 + 1
}
