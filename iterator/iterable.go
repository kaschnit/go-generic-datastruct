package iterator

type ForwardIterable[K any, V any] interface {
	Iterator() (iter ForwardIterator[K, V], ok bool)
	IteratorReverse() (iter ForwardIterator[K, V], ok bool)
}
