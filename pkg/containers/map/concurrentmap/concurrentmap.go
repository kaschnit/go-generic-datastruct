package concurrentmap

import (
	"strings"
	"sync"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/enumerable/abort"
	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
)

func MakeThreadSafe[K any, V any](m mapp.Map[K, V]) *ConcurrentMap[K, V] {
	if c, ok := m.(*ConcurrentMap[K, V]); ok {
		return c
	}

	return &ConcurrentMap[K, V]{
		inner:  m,
		rwlock: sync.RWMutex{},
	}
}

type ConcurrentMap[K any, V any] struct {
	inner  mapp.Map[K, V]
	rwlock sync.RWMutex
}

func (m *ConcurrentMap[K, V]) Empty() bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Empty()
}

func (m *ConcurrentMap[K, V]) Size() int {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Size()
}

func (m *ConcurrentMap[K, V]) Clear() {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.inner.Clear()
}

func (m *ConcurrentMap[K, V]) String() string {
	sb := strings.Builder{}
	sb.WriteString("[Concurrent]")

	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	sb.WriteString(m.inner.String())

	return sb.String()
}

func (m *ConcurrentMap[K, V]) ForEach(op enumerable.Op[K, V]) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	m.inner.ForEach(op)
}

func (m *ConcurrentMap[K, V]) Any(predicate enumerable.Predicate[K, V]) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Any(predicate)
}

func (m *ConcurrentMap[K, V]) All(predicate enumerable.Predicate[K, V]) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.All(predicate)
}

func (m *ConcurrentMap[K, V]) Find(predicate enumerable.Predicate[K, V]) (key K, value V, ok bool) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Find(predicate)
}

func (m *ConcurrentMap[K, V]) Keys(signal abort.Signal) <-chan K {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Keys(signal)
}

func (m *ConcurrentMap[K, V]) Values(signal abort.Signal) <-chan V {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Values(signal)
}

func (m *ConcurrentMap[K, V]) Items(signal abort.Signal) <-chan enumerable.KeyValue[K, V] {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Items(signal)
}

func (m *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.Get(key)
}

func (m *ConcurrentMap[K, V]) Put(key K, value V) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.inner.Put(key, value)
}

func (m *ConcurrentMap[K, V]) PutAll(entries ...entry.Entry[K, V]) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.inner.PutAll(entries...)
}

func (m *ConcurrentMap[K, V]) RemoveKey(key K) bool {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	return m.inner.RemoveKey(key)
}

func (m *ConcurrentMap[K, V]) RemoveAllKeys(keys ...K) int {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	return m.inner.RemoveAllKeys(keys...)
}

func (m *ConcurrentMap[K, V]) ContainsKey(key K) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.ContainsKey(key)
}

func (m *ConcurrentMap[K, V]) ContainsAllKeys(keys ...K) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.ContainsAllKeys(keys...)
}

func (m *ConcurrentMap[K, V]) ContainsAnyKey(keys ...K) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.inner.ContainsAnyKey(keys...)
}
