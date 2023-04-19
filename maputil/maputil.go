package maputil

import "github.com/kaschnit/go-ds/enumerable"

func ForEach[K comparable, V any](m map[K]V, op enumerable.Op[K, V]) {
	for key, value := range m {
		op(key, value)
	}
}

func Filter[K comparable, V any](m map[K]V, predicate enumerable.Predicate[K, V]) map[K]V {
	result := make(map[K]V, 0)
	for key, value := range m {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return result
}

func All[K comparable, V any](m map[K]V, predicate enumerable.Predicate[K, V]) bool {
	for key, value := range m {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

func Any[K comparable, V any](m map[K]V, predicate enumerable.Predicate[K, V]) bool {
	for key, value := range m {
		if predicate(key, value) {
			return true
		}
	}
	return false
}

func Find[K comparable, V any](m map[K]V, predicate enumerable.Predicate[K, V]) (key K, result V, ok bool) {
	for key, value := range m {
		if predicate(key, value) {
			return key, value, true
		}
	}
	return *new(K), *new(V), false
}

func Map[K comparable, V any, R any](m map[K]V, mapper enumerable.Mapper[K, V, R]) []R {
	result := make([]R, 0, len(m))
	for key, value := range m {
		result = append(result, mapper(key, value))
	}
	return result
}

func MapMap[K comparable, V any, KR comparable, VR comparable](m map[K]V, mapper enumerable.MapMapper[K, V, KR, VR]) map[KR]VR {
	result := make(map[KR]VR, 0)
	for key, value := range m {
		mappedKey, mappedValue := mapper(key, value)
		result[mappedKey] = mappedValue
	}
	return result
}
