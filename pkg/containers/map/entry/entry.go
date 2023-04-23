package entry

import "fmt"

type Entry[K any, V any] struct {
	key   K
	value V
}

func New[K any, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{
		key:   key,
		value: value,
	}
}

func NewFromMap[K comparable, V any](m map[K]V) []Entry[K, V] {
	result := make([]Entry[K, V], 0, len(m))
	for key, value := range m {
		result = append(result, Entry[K, V]{key: key, value: value})
	}
	return result
}

func NewRef[K any, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{
		key:   key,
		value: value,
	}
}

func (e *Entry[K, V]) Key() K {
	return e.key
}

func (e *Entry[K, V]) Value() V {
	return e.value
}

func (e *Entry[K, V]) String() string {
	return fmt.Sprintf("Entry{Key:%v, Value:%v}", e.key, e.value)
}
