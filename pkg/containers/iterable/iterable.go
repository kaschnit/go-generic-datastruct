package iterable

import "github.com/kaschnit/go-ds/pkg/iterator"

type ForwardIterable[K any, V any] interface {
	Iterator() (iter iterator.ForwardIterator[K, V], ok bool)
}

type ReverseIterable[K any, V any] interface {
	IteratorReverse() (iter iterator.ForwardIterator[K, V], ok bool)
}
