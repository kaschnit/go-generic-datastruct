package concurrentqueue

import (
	"strings"
	"sync"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
)

func MakeThreadSafe[T any](q queue.Queue[T]) *ConcurrentQueue[T] {
	if c, ok := q.(*ConcurrentQueue[T]); ok {
		return c
	}

	return &ConcurrentQueue[T]{
		inner:  q,
		rwlock: sync.RWMutex{},
	}
}

type ConcurrentQueue[T any] struct {
	inner  queue.Queue[T]
	rwlock sync.RWMutex
}

func (q *ConcurrentQueue[T]) Empty() bool {
	q.rwlock.RLock()
	defer q.rwlock.RUnlock()

	return q.inner.Empty()
}

func (q *ConcurrentQueue[T]) Size() int {
	q.rwlock.RLock()
	defer q.rwlock.RUnlock()

	return q.inner.Size()
}

func (q *ConcurrentQueue[T]) Clear() {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()

	q.inner.Clear()
}

func (q *ConcurrentQueue[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[Concurrent]")

	q.rwlock.RLock()
	defer q.rwlock.RUnlock()
	sb.WriteString(q.inner.String())

	return sb.String()
}

func (q *ConcurrentQueue[T]) Push(value T) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()

	q.inner.Push(value)
}

func (q *ConcurrentQueue[T]) PushAll(values ...T) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()

	q.inner.PushAll(values...)
}

func (q *ConcurrentQueue[T]) Pop() (value T, ok bool) {
	q.rwlock.Lock()
	defer q.rwlock.Unlock()

	return q.inner.Pop()
}

func (q *ConcurrentQueue[T]) Peek() (value T, ok bool) {
	q.rwlock.RLock()
	defer q.rwlock.RUnlock()

	return q.inner.Peek()
}
