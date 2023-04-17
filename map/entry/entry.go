package entry

type Entry[K any, V any] struct {
	key   K
	value V
}

func New[K any, V any](key K, value V) *Entry[K, V] {
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
