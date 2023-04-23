package hashmap

import (
	"strings"

	"github.com/kaschnit/go-ds/pkg/compare"
	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
)

type HashMapBuilder[K any, HK comparable, V any] struct {
	hashkey compare.HashKey[K, HK]
	entries map[HK]entry.Entry[K, V]
}

func NewBuilder[K any, HK comparable, V any](hashkey compare.HashKey[K, HK]) *HashMapBuilder[K, HK, V] {
	return &HashMapBuilder[K, HK, V]{
		hashkey: hashkey,
		entries: make(map[HK]entry.Entry[K, V]),
	}
}

func (b *HashMapBuilder[K, HK, V]) Put(key K, value V) *HashMapBuilder[K, HK, V] {
	b.entries[b.hashkey(key)] = entry.New(key, value)

	return b
}

func (b *HashMapBuilder[K, HK, V]) PutAll(entries ...entry.Entry[K, V]) *HashMapBuilder[K, HK, V] {
	for _, entry := range entries {
		b.entries[b.hashkey(entry.Key())] = entry
	}

	return b
}

func (b *HashMapBuilder[K, HK, V]) Build() *HashMap[K, HK, V] {
	return &HashMap[K, HK, V]{
		hashkey: b.hashkey,
		entries: b.entries,
	}
}

type HashMap[K any, HK comparable, V any] struct {
	hashkey compare.HashKey[K, HK]
	entries map[HK]entry.Entry[K, V]
}

func New[K comparable, V any](entries ...entry.Entry[K, V]) *HashMap[K, K, V] {
	return NewBuilder[K, K, V](compare.IdentityHashKey[K]).PutAll(entries...).Build()
}

func (m *HashMap[K, HK, V]) Empty() bool {
	return m.Size() == 0
}

func (m *HashMap[K, HK, V]) Size() int {
	return len(m.entries)
}

func (m *HashMap[K, HK, V]) Clear() {
	for k := range m.entries {
		delete(m.entries, k)
	}
}

func (m *HashMap[K, HK, V]) String() string {
	sb := strings.Builder{}
	sb.WriteString("HashMap\n")
	strs := []string{}
	for _, entry := range m.entries {
		strs = append(strs, entry.String())
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (m *HashMap[K, HK, V]) ForEach(op enumerable.Op[K, V]) {
	for _, entry := range m.entries {
		op(entry.Key(), entry.Value())
	}
}

func (m *HashMap[K, HK, V]) Any(predicate enumerable.Predicate[K, V]) bool {
	for _, entry := range m.entries {
		if predicate(entry.Key(), entry.Value()) {
			return true
		}
	}
	return false
}

func (m *HashMap[K, HK, V]) All(predicate enumerable.Predicate[K, V]) bool {
	for _, entry := range m.entries {
		if !predicate(entry.Key(), entry.Value()) {
			return false
		}
	}
	return true
}

func (m *HashMap[K, HK, V]) Find(predicate enumerable.Predicate[K, V]) (key K, value V, ok bool) {
	for _, entry := range m.entries {
		if predicate(key, value) {
			return entry.Key(), entry.Value(), true
		}
	}
	return *new(K), *new(V), false
}

func (m *HashMap[K, HK, V]) Keys(abort <-chan struct{}) <-chan K {
	ch := make(chan K, 1)
	go func() {
		defer close(ch)
		for _, entry := range m.entries {
			select {
			case ch <- entry.Key():
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, HK, V]) Values(abort <-chan struct{}) <-chan V {
	ch := make(chan V, 1)
	go func() {
		defer close(ch)
		for _, entry := range m.entries {
			select {
			case ch <- entry.Value():
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, HK, V]) Items(abort <-chan struct{}) <-chan enumerable.KeyValue[K, V] {
	ch := make(chan enumerable.KeyValue[K, V], 1)
	go func() {
		defer close(ch)
		for _, entry := range m.entries {
			select {
			case ch <- enumerable.KeyValue[K, V]{
				Key:   entry.Key(),
				Value: entry.Value(),
			}:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, HK, V]) Put(key K, value V) {
	m.entries[m.hashkey(key)] = entry.New(key, value)
}

func (m *HashMap[K, HK, V]) PutAll(entries ...entry.Entry[K, V]) {
	for _, entry := range entries {
		m.Put(entry.Key(), entry.Value())
	}
}

func (m *HashMap[K, HK, V]) RemoveKey(key K) bool {
	hashedKey := m.hashkey(key)
	contained := m.containsHashedKey(hashedKey)
	delete(m.entries, hashedKey)
	return contained
}

func (m *HashMap[K, HK, V]) RemoveAllKeys(keys ...K) bool {
	contained := true
	for _, key := range keys {
		contained = contained && m.RemoveKey(key)
	}
	return contained
}

func (m *HashMap[K, HK, V]) ContainsKey(key K) bool {
	return m.containsHashedKey(m.hashkey(key))
}

func (m *HashMap[K, HK, V]) ContainsAllKeys(keys ...K) bool {
	for _, key := range keys {
		if !m.ContainsKey(key) {
			return false
		}
	}
	return true
}

func (m *HashMap[K, HK, V]) ContainsAnyKey(keys ...K) bool {
	for _, key := range keys {
		if m.ContainsKey(key) {
			return true
		}
	}
	return false
}

func (m *HashMap[K, HK, V]) containsHashedKey(key HK) bool {
	_, contains := m.entries[key]
	return contains
}
