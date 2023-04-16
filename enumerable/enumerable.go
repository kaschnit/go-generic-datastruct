package enumerable

type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

type Op[K any, V any] func(key K, value V)
type Predicate[K any, V any] func(key K, value V) bool

type Enumerable[K any, V any] interface {
	ForEach(op Op[K, V])
	Any(predicate Predicate[K, V]) bool
	All(predicate Predicate[K, V]) bool
	Find(predicate Predicate[K, V]) (key K, value V, ok bool)
	Keys(abort <-chan struct{}) <-chan K
	Values(abort <-chan struct{}) <-chan V
	Items(abort <-chan struct{}) <-chan KeyValue[K, V]
}
