package container

type Container interface {
	Empty() bool
	Size() int
	Clear()
	String() string
}
