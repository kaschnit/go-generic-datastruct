package concurrentset

import (
	"strings"
	"sync"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/enumerable/abort"
	"github.com/kaschnit/go-ds/pkg/containers/set"
)

func MakeThreadSafe[T comparable](s set.Set[T]) *ConcurrentSet[T] {
	if c, ok := s.(*ConcurrentSet[T]); ok {
		return c
	}

	return &ConcurrentSet[T]{
		inner:  s,
		rwlock: sync.RWMutex{},
	}
}

type ConcurrentSet[T comparable] struct {
	inner  set.Set[T]
	rwlock sync.RWMutex
}

func (s *ConcurrentSet[T]) Empty() bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Empty()
}

func (s *ConcurrentSet[T]) Size() int {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Size()
}

func (s *ConcurrentSet[T]) Clear() {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	s.inner.Clear()
}

func (s *ConcurrentSet[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[Concurrent]")

	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	sb.WriteString(s.inner.String())

	return sb.String()
}

func (s *ConcurrentSet[T]) ForEach(op enumerable.Op[T, T]) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	s.inner.ForEach(op)
}

func (s *ConcurrentSet[T]) Any(predicate enumerable.Predicate[T, T]) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Any(predicate)
}

func (s *ConcurrentSet[T]) All(predicate enumerable.Predicate[T, T]) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.All(predicate)
}

func (s *ConcurrentSet[T]) Find(predicate enumerable.Predicate[T, T]) (key T, value T, ok bool) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Find(predicate)
}

func (s *ConcurrentSet[T]) Keys(signal abort.Signal) <-chan T {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Keys(signal)
}

func (s *ConcurrentSet[T]) Values(signal abort.Signal) <-chan T {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Values(signal)
}

func (s *ConcurrentSet[T]) Items(signal abort.Signal) <-chan enumerable.KeyValue[T, T] {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Items(signal)
}

func (s *ConcurrentSet[T]) Add(value T) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	s.inner.Add(value)
}

func (s *ConcurrentSet[T]) AddAll(values ...T) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	s.inner.AddAll(values...)
}

func (s *ConcurrentSet[T]) Contains(value T) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Contains(value)
}

func (s *ConcurrentSet[T]) Remove(value T) bool {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	return s.inner.Remove(value)
}

func (s *ConcurrentSet[T]) RemoveAll(values ...T) int {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	return s.inner.RemoveAll(values...)
}

func (s *ConcurrentSet[T]) ContainsAll(values ...T) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.ContainsAll(values...)
}

func (s *ConcurrentSet[T]) ContainsAny(values ...T) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.ContainsAny(values...)
}
