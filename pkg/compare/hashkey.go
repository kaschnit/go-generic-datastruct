package compare

type HashKey[K any, H comparable] func(key K) H

func IdentityHashKey[K comparable](key K) K {
	return key
}
