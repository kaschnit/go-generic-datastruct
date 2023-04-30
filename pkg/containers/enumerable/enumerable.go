package enumerable

import (
	"github.com/kaschnit/go-ds/pkg/containers/container"
	"github.com/kaschnit/go-ds/pkg/containers/enumerable/abort"
)

// KeyValue is a pair of a key and a value
type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

// Op is an arbitrary operation that does not return a value.
type Op[K any, V any] func(key K, value V)

// Predicate defines a condition and should return true if the condition holds.
type Predicate[K any, V any] func(key K, value V) bool

// Mapper is a function that maps a single input to an output.
type Mapper[K any, V any, R any] func(key K, value V) R

// MapMapper is a function that maps a single input to a pair of outputs, which
// are used to construct a key-value map.
type MapMapper[K comparable, V any, KR comparable, VR any] func(key K, value V) (KR, VR)

// Enumerable defines an interface for all containers that can be "enumerated".
type Enumerable[K any, V any] interface {
	ForEach(op Op[K, V])
	Any(predicate Predicate[K, V]) bool
	All(predicate Predicate[K, V]) bool
	Find(predicate Predicate[K, V]) (key K, value V, ok bool)
	Keys(signal abort.Signal) <-chan K
	Values(signal abort.Signal) <-chan V
	Items(signal abort.Signal) <-chan KeyValue[K, V]
}

type enumerableContainer[K any, V any] interface {
	container.Container
	Enumerable[K, V]
}

// Map iterates over any enumerable container and applies a transformation to each item
// in the container, producing a slice of the transformed items.
// This does not mutate the container, although it's possible for the mapper function to
// mutate the items in the container.
func Map[K any, V any, R any](e enumerableContainer[K, V], mapper Mapper[K, V, R]) []R {
	result := make([]R, 0, e.Size())
	e.ForEach(func(key K, value V) {
		result = append(result, mapper(key, value))
	})
	return result
}

// Map iterates over any enumerable container and applies a transformation to each item
// in the container, producing a map of the transformed items.
// This does not mutate the container, although it's possible for the mapper function to
// mutate the items in the container.
func MapMap[K comparable, V any, KR comparable, VR comparable](e enumerableContainer[K, V], mapper MapMapper[K, V, KR, VR]) map[KR]VR {
	result := make(map[KR]VR, 0)
	e.ForEach(func(key K, value V) {
		mappedKey, mappedValue := mapper(key, value)
		result[mappedKey] = mappedValue
	})
	return result
}
