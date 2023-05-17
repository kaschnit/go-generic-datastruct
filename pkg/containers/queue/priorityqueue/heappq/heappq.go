package heappq

import (
	"fmt"
	"strings"

	compare "github.com/kaschnit/go-ds/pkg/compare"
	"golang.org/x/exp/constraints"
)

type Builder[T any] struct {
	comparator compare.Comparator[T]
	items      []T
}

func NewBuilder[T any](comparator compare.Comparator[T]) *Builder[T] {
	return &Builder[T]{
		comparator: comparator,
	}
}

func (b *Builder[T]) AddItems(items ...T) *Builder[T] {
	b.items = append(b.items, items...)

	return b
}

func (b *Builder[T]) Build() *HeapPQ[T] {
	q := HeapPQ[T]{
		comparator: b.comparator,
		items:      make([]T, 1),
	}
	q.PushAll(b.items...)

	return &q
}

type HeapPQ[T any] struct {
	comparator compare.Comparator[T]
	items      []T
}

func New[T constraints.Ordered](values ...T) *HeapPQ[T] {
	return NewBuilder(compare.OrderedComparator[T]).AddItems(values...).Build()
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
	qCpy := q.Copy()
	sb := strings.Builder{}
	sb.WriteString("HeapPQ\n")

	strs := make([]string, 0, qCpy.Size())
	for item, ok := qCpy.Pop(); ok; item, ok = qCpy.Pop() {
		strs = append(strs, fmt.Sprintf("%v", item))
	}

	sb.WriteString(strings.Join(strs, ","))

	return sb.String()
}

func (q *HeapPQ[T]) Push(value T) {
	// Push onto the end
	q.items = append(q.items, value)

	// Fix the heap invariant
	q.percolateUp()
}

func (q *HeapPQ[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *HeapPQ[T]) Pop() (T, bool) {
	value, ok := q.Peek()
	if !ok {
		return value, false
	}

	// Move the last item to the root
	q.items[1], q.items[len(q.items)-1] = q.items[len(q.items)-1], q.items[1]

	// Remove the last item, which was the root
	q.items = q.items[:len(q.items)-1]

	// Fix the heap invariant
	q.percolateDown()

	return value, true
}

func (q *HeapPQ[T]) Peek() (T, bool) {
	if q.Empty() {
		return *new(T), false
	}

	return q.items[1], true
}

func (q *HeapPQ[T]) Copy() *HeapPQ[T] {
	return NewBuilder(q.comparator).AddItems(q.items[1:]...).Build()
}

func (q *HeapPQ[T]) percolateUp() {
	// Percolate up to maintain heap invariant
	for fixIndex := len(q.items) - 1; fixIndex != 1; {
		parentIndex := parent(fixIndex)

		// Check whether the parent is higher priority
		isParentInPlace := q.comparator(q.items[fixIndex], q.items[parentIndex]) != compare.PriorityLeftHigher

		// If the parent priority is higher, we're done percolating up.
		if isParentInPlace {
			break
		}

		// Swap the current index with the parent and continue
		q.items[fixIndex], q.items[parentIndex] = q.items[parentIndex], q.items[fixIndex]
		fixIndex = parentIndex
	}
}

func (q *HeapPQ[T]) percolateDown() {
	// Percolate down to maintain the heap invariant
	for fixIndex := 1; fixIndex != len(q.items)-1; {
		leftIndex := leftChild(fixIndex)
		rightIndex := rightChild(fixIndex)

		childIndex := -1

		if leftIndex < len(q.items) && rightIndex < len(q.items) {
			cmp := q.comparator(q.items[leftIndex], q.items[rightIndex])
			if cmp == compare.PriorityLeftHigher {
				childIndex = leftIndex
			} else {
				childIndex = rightIndex
			}
		} else if leftIndex < len(q.items) {
			childIndex = leftIndex
		} else if rightIndex < len(q.items) {
			childIndex = rightIndex
		}

		if childIndex >= 0 {
			childInPlace := q.comparator(q.items[childIndex], q.items[fixIndex]) != compare.PriorityLeftHigher

			// Swap with the child if the child is not in place.
			if !childInPlace {
				q.items[fixIndex], q.items[childIndex] = q.items[childIndex], q.items[fixIndex]
				fixIndex = childIndex

				continue
			}
		}

		// Node is in place, done.
		break
	}
}

//nolint:gomnd
func parent(index int) int {
	return index / 2
}

//nolint:gomnd
func leftChild(index int) int {
	return index * 2
}

func rightChild(index int) int {
	return index*2 + 1
}
