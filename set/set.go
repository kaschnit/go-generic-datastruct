package set

import "github.com/kaschnit/go-ds/container"

type Set[T any] interface {
	container.Container[T]

	Add(value T)
	AddAll(values ...T)
	Remove(value T) bool
	RemoveAll(values ...T) bool
	Contains(value T) bool
	ContainsAll(values ...T) bool
	ContainsAny(values ...T) bool
}
