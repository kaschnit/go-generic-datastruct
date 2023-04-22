package enumerable

import "github.com/kaschnit/go-ds/pkg/containers/container"

type KeyValue[K any, V any] struct {
	Key   K
	Value V
}

type Op[K any, V any] func(key K, value V)
type Predicate[K any, V any] func(key K, value V) bool
type Mapper[K any, V any, R any] func(key K, value V) R
type MapMapper[K comparable, V any, KR comparable, VR any] func(key K, value V) (KR, VR)

type Enumerable[K any, V any] interface {
	ForEach(op Op[K, V])
	Any(predicate Predicate[K, V]) bool
	All(predicate Predicate[K, V]) bool
	Find(predicate Predicate[K, V]) (key K, value V, ok bool)
	Keys(abort <-chan struct{}) <-chan K
	Values(abort <-chan struct{}) <-chan V
	Items(abort <-chan struct{}) <-chan KeyValue[K, V]
}

type enumerableContainer[K any, V any] interface {
	container.Container
	Enumerable[K, V]
}

func Map[K any, V any, R any](e enumerableContainer[K, V], mapper Mapper[K, V, R]) []R {
	result := make([]R, 0, e.Size())
	e.ForEach(func(key K, value V) {
		result = append(result, mapper(key, value))
	})
	return result
}

func MapMap[K comparable, V any, KR comparable, VR comparable](e enumerableContainer[K, V], mapper MapMapper[K, V, KR, VR]) map[KR]VR {
	result := make(map[KR]VR, 0)
	e.ForEach(func(key K, value V) {
		mappedKey, mappedValue := mapper(key, value)
		result[mappedKey] = mappedValue
	})
	return result
}
