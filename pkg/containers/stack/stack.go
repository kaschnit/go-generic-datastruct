package stack

import "github.com/kaschnit/go-ds/pkg/containers/container"

type Stack[T any] interface {
	container.Container

	Push(value T)
	PushAll(values ...T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
