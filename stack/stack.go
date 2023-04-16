package stack

import "github.com/kaschnit/go-ds/container"

type Stack[T any] interface {
	container.Container[T]

	Push(value T)
	PushAll(values ...T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)
}
