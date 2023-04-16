package arraystack

import (
	"fmt"
	"strings"
)

type ArrayStack[T any] struct {
	values []T
}

func New[T any](values ...T) *ArrayStack[T] {
	stack := &ArrayStack[T]{
		values: make([]T, len(values)),
	}
	copy(stack.values, values)
	return stack
}

func (s *ArrayStack[T]) Empty() bool {
	return s.Size() == 0
}

func (s *ArrayStack[T]) Size() int {
	return len(s.values)
}

func (s *ArrayStack[T]) Clear() {
	s.values = make([]T, 0)
}

func (s *ArrayStack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("ArrayStack\n")
	strs := make([]string, s.Size())
	for i, value := range s.values {
		strs[i] = fmt.Sprintf("%v", value)
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (s *ArrayStack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *ArrayStack[T]) PushAll(values ...T) {
	for _, value := range values {
		s.Push(value)
	}
}

func (s *ArrayStack[T]) Pop() (value T, ok bool) {
	value, ok = s.Peek()
	if ok {
		s.values = s.values[:len(s.values)-1]
	}
	return value, ok
}

func (s *ArrayStack[T]) Peek() (value T, ok bool) {
	if s.Empty() {
		return *new(T), false
	}
	return s.values[len(s.values)-1], true
}
