package blockingqueue

import (
	"context"
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
	"golang.org/x/sync/semaphore"
)

type ContextProvider func() context.Context

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
		sem:        *semaphore.NewWeighted(int64(b.bufSize)),
		bufSize:    b.bufSize,
	}
	q.PushAll(b.items...)

	return q
}

type BlockingQueue[T any] struct {
	linkedList  *linkedlist.DoubleLinkedList[T]
	bufSize     int
	sem         semaphore.Weighted
	ctxProvider ContextProvider
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
	q.PushWithContext(context.Background(), value)
}

func (q *BlockingQueue[T]) PushWithContext(ctx context.Context, value T) {
	q.sem.Acquire(ctx, 1)
	q.linkedList.Prepend(value)
}

func (q *BlockingQueue[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *BlockingQueue[T]) Pop() (value T, ok bool) {
	val, ok := q.linkedList.PopBack()
	if ok {
		q.sem.Release(1)
	}
	return val, ok
}

func (q *BlockingQueue[T]) Peek() (value T, ok bool) {
	return q.linkedList.GetBack()
}
