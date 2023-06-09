package concurrentset

import (
	"strings"
	"sync"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/set"
)

func MakeThreadSafe[T comparable](otherSet set.Set[T]) *ConcurrentSet[T] {
	if c, ok := otherSet.(*ConcurrentSet[T]); ok {
		return c
	}

	return &ConcurrentSet[T]{
		inner:  otherSet,
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
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("[Concurrent]")

	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	stringBuilder.WriteString(s.inner.String())

	return stringBuilder.String()
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

func (s *ConcurrentSet[T]) Find(predicate enumerable.Predicate[T, T]) (T, T, bool) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	return s.inner.Find(predicate)
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
