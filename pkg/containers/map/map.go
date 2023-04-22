package mapp

import (
	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
)

type Map[K any, V any] interface {
	enumerable.Enumerable[K, V]

	Put(key K, value V)
	PutAll(entries ...entry.Entry[K, V])
	RemoveKey(key K) bool
	RemoveAllKeys(keys ...K) bool
	ContainsKey(key K) bool
	ContainsAllKeys(keys ...K) bool
	ContainsAnyKey(keys ...K) bool
}
