package hashmap

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
)

type HashMap[K comparable, V any] struct {
	values map[K]V
}

func New[K comparable, V any](values ...entry.Entry[K, V]) *HashMap[K, V] {
	m := HashMap[K, V]{
		values: make(map[K]V),
	}
	m.PutAll(values...)
	return &m
}

func (m *HashMap[K, V]) Empty() bool {
	return m.Size() == 0
}

func (m *HashMap[K, V]) Size() int {
	return len(m.values)
}

func (m *HashMap[K, V]) Clear() {
	for k := range m.values {
		delete(m.values, k)
	}
}

func (m *HashMap[K, V]) String() string {
	sb := strings.Builder{}
	sb.WriteString("HashSet\n")
	strs := []string{}
	for k := range m.values {
		strs = append(strs, fmt.Sprintf("%v", k))
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (m *HashMap[K, V]) ForEach(op enumerable.Op[K, V]) {
	for key, value := range m.values {
		op(key, value)
	}
}

func (m *HashMap[K, V]) Any(predicate enumerable.Predicate[K, V]) bool {
	for key, value := range m.values {
		if predicate(key, value) {
			return true
		}
	}
	return false
}

func (m *HashMap[K, V]) All(predicate enumerable.Predicate[K, V]) bool {
	for key, value := range m.values {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

func (m *HashMap[K, V]) Find(predicate enumerable.Predicate[K, V]) (key K, value V, ok bool) {
	for key, value := range m.values {
		if predicate(key, value) {
			return key, value, true
		}
	}
	return *new(K), *new(V), false
}

func (m *HashMap[K, V]) Keys(abort <-chan struct{}) <-chan K {
	ch := make(chan K, 1)
	go func() {
		defer close(ch)
		for key := range m.values {
			select {
			case ch <- key:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, V]) Values(abort <-chan struct{}) <-chan V {
	ch := make(chan V, 1)
	go func() {
		defer close(ch)
		for _, value := range m.values {
			select {
			case ch <- value:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, V]) Items(abort <-chan struct{}) <-chan enumerable.KeyValue[K, V] {
	ch := make(chan enumerable.KeyValue[K, V], 1)
	go func() {
		defer close(ch)
		for key, value := range m.values {
			select {
			case ch <- enumerable.KeyValue[K, V]{
				Key:   key,
				Value: value,
			}:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (m *HashMap[K, V]) Put(key K, value V) {
	m.values[key] = value
}

func (m *HashMap[K, V]) PutAll(entries ...entry.Entry[K, V]) {
	for _, entry := range entries {
		m.Put(entry.Key(), entry.Value())
	}
}

func (m *HashMap[K, V]) RemoveKey(key K) bool {
	contained := m.ContainsKey(key)
	delete(m.values, key)
	return contained
}

func (m *HashMap[K, V]) RemoveAllKeys(keys ...K) bool {
	contained := true
	for _, key := range keys {
		contained = contained && m.RemoveKey(key)
	}
	return contained
}

func (m *HashMap[K, V]) ContainsKey(key K) bool {
	_, contains := m.values[key]
	return contains
}

func (m *HashMap[K, V]) ContainsAllKeys(keys ...K) bool {
	for _, key := range keys {
		if !m.ContainsKey(key) {
			return false
		}
	}
	return true
}

func (m *HashMap[K, V]) ContainsAnyKey(keys ...K) bool {
	for _, key := range keys {
		if m.ContainsKey(key) {
			return true
		}
	}
	return false
}
