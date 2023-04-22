package iterator

type iterator[K any, V any] interface {
	Key() (key K, ok bool)
	Value() (value V, ok bool)
}

type ForwardIterator[K any, V any] interface {
	iterator[K, V]

	Next() (next ForwardIterator[K, V], ok bool)
	HasNext() bool
}
