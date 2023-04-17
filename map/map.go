package mapp

import (
	"github.com/kaschnit/go-ds/container"
	"github.com/kaschnit/go-ds/enumerable"
)

type Map[K any, V any] interface {
	container.Container
	enumerable.Enumerable[K, V]

	Put(key K, value V)
	PutAll(values ...Entry[K, V])
	Remove(key K) bool
	RemoveAll(keys ...K) bool
	Contains(entry Entry[K, V]) bool
	ContainsAll(entries ...Entry[K, V]) bool
	ContainsAny(entries ...Entry[K, V]) bool
	ContainsKey(key K) bool
	ContainsAllKeys(keys ...K) bool
	ContainsAnyKey(keys ...K) bool
	ContainsValue(value V) bool
	ContainsAllValues(values ...V) bool
	ContainsAnyValues(values ...V) bool
}

type Entry[K any, V any] struct {
	key   K
	value V
}

func NewEntry[K any, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{
		key:   key,
		value: value,
	}
}
