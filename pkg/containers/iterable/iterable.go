package iterable

import "github.com/kaschnit/go-ds/pkg/iterator"

// ForwardIterable is a type that can produce an iterator.
type ForwardIterable[K any, V any] interface {
	Iterator() (iter iterator.ForwardIterator[K, V], ok bool)
}

// ReverseIterable is a type that supports ordered iteration, and can produce an iterator
// to iterate in the reverse direction of the container's order.
type ReverseIterable[K any, V any] interface {
	IteratorReverse() (iter iterator.ForwardIterator[K, V], ok bool)
}
