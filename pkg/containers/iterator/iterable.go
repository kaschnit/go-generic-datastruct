package iterator

type ForwardIterable[K any, V any] interface {
	Iterator() (iter ForwardIterator[K, V], ok bool)
}

type ReverseIterable[K any, V any] interface {
	IteratorReverse() (iter ForwardIterator[K, V], ok bool)
}
