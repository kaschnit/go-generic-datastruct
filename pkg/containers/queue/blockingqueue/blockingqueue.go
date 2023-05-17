package blockingqueue

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
)

type Builder[T any] struct {
	bufSize int
	items   []T
}

func NewBuilder[T any](bufSize int) *Builder[T] {
	return &Builder[T]{
		bufSize: bufSize,
	}
}

func (b *Builder[T]) AddItems(items ...T) *Builder[T] {
	b.items = append(b.items, items...)
	if len(b.items) > b.bufSize {
		b.items = b.items[:b.bufSize]
	}

	return b
}

func (b *Builder[T]) Build() *BlockingQueue[T] {
	q := &BlockingQueue[T]{
		linkedList: linkedlist.NewDoubleLinked[T](),
		sem:        make(chan struct{}, b.bufSize),
		bufSize:    b.bufSize,
	}
	q.PushAll(b.items...)

	return q
}

type BlockingQueue[T any] struct {
	linkedList *linkedlist.DoubleLinkedList[T]
	bufSize    int
	sem        chan struct{}
}

func (q *BlockingQueue[T]) Empty() bool {
	return q.linkedList.Empty()
}

func (q *BlockingQueue[T]) Size() int {
	return q.linkedList.Size()
}

func (q *BlockingQueue[T]) Clear() {
	q.linkedList.Clear()
}

func (q *BlockingQueue[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("BlockingQueue[capacity=%d]\n", q.bufSize))

	strs := make([]string, 0, q.Size())
	q.linkedList.ForEach(func(_ int, value T) {
		strs = append(strs, fmt.Sprintf("%v", value))
	})

	sb.WriteString(strings.Join(strs, ","))

	return sb.String()
}

func (q *BlockingQueue[T]) Push(value T) {
	q.sem <- struct{}{}
	q.linkedList.Prepend(value)
}

func (q *BlockingQueue[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *BlockingQueue[T]) Pop() (T, bool) {
	<-q.sem

	return q.linkedList.PopBack()
}

func (q *BlockingQueue[T]) Peek() (T, bool) {
	return q.linkedList.GetBack()
}
