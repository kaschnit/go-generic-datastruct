package iterator

type iterator[K any, V any] interface {
	Key() (key K, ok bool)
	Value() (value V, ok bool)
}

// ForwardIterator can be used to iterate in one direction over a collection of items.
// This does not necessarily imply consistent ordering or iteration.
type ForwardIterator[K any, V any] interface {
	iterator[K, V]

	Next() (next ForwardIterator[K, V], ok bool)
	HasNext() bool
}
