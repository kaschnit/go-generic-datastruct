package linkedqueue

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
)

type LinkedQueue[T any] struct {
	linkedList *linkedlist.DoubleLinkedList[T]
}

func New[T any](values ...T) *LinkedQueue[T] {
	q := LinkedQueue[T]{
		linkedList: linkedlist.NewDoubleLinked[T](),
	}
	q.PushAll(values...)
	return &q
}

func (q *LinkedQueue[T]) Empty() bool {
	return q.linkedList.Empty()
}

func (q *LinkedQueue[T]) Size() int {
	return q.linkedList.Size()
}

func (q *LinkedQueue[T]) Clear() {
	q.linkedList.Clear()
}

func (q *LinkedQueue[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("LinkedQueue\n")

	strs := make([]string, 0, q.Size())
	q.linkedList.ForEach(func(_ int, value T) {
		strs = append(strs, fmt.Sprintf("%v", value))
	})
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (q *LinkedQueue[T]) Push(value T) {
	q.linkedList.Prepend(value)
}

func (q *LinkedQueue[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *LinkedQueue[T]) Pop() (value T, ok bool) {
	return q.linkedList.PopBack()
}

func (q *LinkedQueue[T]) Peek() (value T, ok bool) {
	return q.linkedList.GetBack()
}
