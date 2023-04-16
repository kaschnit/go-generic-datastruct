package container

type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	String() string
}
