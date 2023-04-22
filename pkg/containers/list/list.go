package list

import (
	"github.com/kaschnit/go-ds/pkg/containers/container"
	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/iterator"
)

type List[T any] interface {
	container.Container
	enumerable.Enumerable[int, T]
	iterator.ForwardIterable[int, T]

	Append(value T)
	AppendAll(values ...T)
	Prepend(value T)
	PrependAll(values ...T)
	Insert(index int, value T) (ok bool)
	InsertAll(index int, values ...T) (ok bool)
	PopBack() (value T, ok bool)
	PopFront() (value T, ok bool)
	GetFront() (value T, ok bool)
	GetBack() (value T, ok bool)
	Get(index int) (value T, ok bool)
}
