package list

import "github.com/kaschnit/go-ds/container"

type List[T any] interface {
	container.Container[T]
}
