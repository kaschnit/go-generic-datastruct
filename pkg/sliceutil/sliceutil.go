package sliceutil

type Op[V any] func(index int, value V)
type Predicate[V any] func(index int, value V) bool
type Mapper[V any, R any] func(index int, value V) R
type MapMapper[V any, KR comparable, VR any] func(key int, value V) (KR, VR)

func ForEach[V any](slice []V, op Op[V]) {
	for i, value := range slice {
		op(i, value)
	}
}

func Filter[V any](slice []V, predicate Predicate[V]) []V {
	result := make([]V, 0)
	for i, value := range slice {
		if predicate(i, value) {
			result = append(result, value)
		}
	}
	return result
}

func All[V any](slice []V, predicate Predicate[V]) bool {
	for i, value := range slice {
		if !predicate(i, value) {
			return false
		}
	}
	return true
}

func Any[V any](slice []V, predicate Predicate[V]) bool {
	for i, value := range slice {
		if predicate(i, value) {
			return true
		}
	}
	return false
}

func Find[V any](slice []V, predicate Predicate[V]) (index int, result V, ok bool) {
	for i, value := range slice {
		if predicate(i, value) {
			return i, value, true
		}
	}
	return -1, *new(V), false
}

func Map[V any, R any](slice []V, mapper Mapper[V, R]) []R {
	result := make([]R, 0, len(slice))
	for i, value := range slice {
		result = append(result, mapper(i, value))
	}
	return result
}

func MapMap[K comparable, V any, KR comparable, VR comparable](slice []V, mapper MapMapper[V, KR, VR]) map[KR]VR {
	result := make(map[KR]VR, 0)
	for index, value := range slice {
		mappedKey, mappedValue := mapper(index, value)
		result[mappedKey] = mappedValue
	}
	return result
}
