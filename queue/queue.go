package queue

import "github.com/kaschnit/go-ds/container"

type Queue[T any] interface {
	container.Container

	Push(value T)
	PushAll(values ...T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
