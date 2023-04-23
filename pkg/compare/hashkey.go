package compare

// HashKey is used to convert a potentially non-comparable value to a comparable value.
// This is useful for computing hash keys for a hash map, for example.
type HashKey[K any, H comparable] func(key K) H

// IdentityHashKey is a no-op on a value that's already comparable value.
// This allows use of the value itself as the hash key with interfaces that require a HashKey.
func IdentityHashKey[K comparable](key K) K {
	return key
}
