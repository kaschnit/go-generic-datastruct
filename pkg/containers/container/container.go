package container

// Container defines a common interface for all containers.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	String() string
}
