package concurrentlist

import (
	"strings"
	"sync"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/iterator"
)

func MakeThreadSafe[T any](l list.List[T]) *ConcurrentList[T] {
	if c, ok := l.(*ConcurrentList[T]); ok {
		return c
	}

	return &ConcurrentList[T]{
		inner:  l,
		rwlock: sync.RWMutex{},
	}
}

type ConcurrentList[T any] struct {
	inner  list.List[T]
	rwlock sync.RWMutex
}

func (l *ConcurrentList[T]) Empty() bool {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.Empty()
}

func (l *ConcurrentList[T]) Size() int {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.Size()
}

func (l *ConcurrentList[T]) Clear() {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()
	l.inner.Clear()
}

func (l *ConcurrentList[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[Concurrent]")

	l.rwlock.RLock()
	defer l.rwlock.RUnlock()
	sb.WriteString(l.inner.String())

	return sb.String()
}

func (l *ConcurrentList[T]) ForEach(op enumerable.Op[int, T]) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	l.inner.ForEach(op)
}

func (l *ConcurrentList[T]) Any(predicate enumerable.Predicate[int, T]) bool {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.Any(predicate)
}

func (l *ConcurrentList[T]) All(predicate enumerable.Predicate[int, T]) bool {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.All(predicate)
}

func (l *ConcurrentList[T]) Find(predicate enumerable.Predicate[int, T]) (int, T, bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.Find(predicate)
}

func (l *ConcurrentList[T]) Iterator() (iterator.ForwardIterator[int, T], bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	// TODO: implement thread-safe iterator.
	// The ConcurrentList.Iterator() method itself is thread-safe, but it's easy
	// to cause data races with the returned iterators.
	return l.inner.Iterator()
}

func (l *ConcurrentList[T]) Append(value T) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	l.inner.Append(value)
}

func (l *ConcurrentList[T]) AppendAll(values ...T) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	l.inner.AppendAll(values...)
}

func (l *ConcurrentList[T]) Prepend(value T) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	l.inner.Prepend(value)
}

func (l *ConcurrentList[T]) PrependAll(values ...T) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	l.inner.PrependAll(values...)
}

func (l *ConcurrentList[T]) Insert(index int, value T) bool {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	return l.inner.Insert(index, value)
}

func (l *ConcurrentList[T]) InsertAll(index int, values ...T) bool {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	return l.inner.InsertAll(index, values...)
}

func (l *ConcurrentList[T]) PopBack() (T, bool) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	return l.inner.PopBack()
}

func (l *ConcurrentList[T]) PopFront() (T, bool) {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	return l.inner.PopFront()
}

func (l *ConcurrentList[T]) GetFront() (T, bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.GetFront()
}

func (l *ConcurrentList[T]) GetBack() (T, bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.GetBack()
}

func (l *ConcurrentList[T]) Get(index int) (T, bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	return l.inner.Get(index)
}
