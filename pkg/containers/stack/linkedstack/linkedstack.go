package linkedstack

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	value T
	prev  *node[T]
}

type LinkedStack[T any] struct {
	head *node[T]
	size int
}

func New[T any](values ...T) *LinkedStack[T] {
	stack := &LinkedStack[T]{
		head: nil,
		size: 0,
	}
	stack.PushAll(values...)

	return stack
}

func (s *LinkedStack[T]) Empty() bool {
	return s.Size() == 0
}

func (s *LinkedStack[T]) Size() int {
	return s.size
}

func (s *LinkedStack[T]) Clear() {
	s.head = nil
	s.size = 0
}

func (s *LinkedStack[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("LinkedStack\n")

	strs := make([]string, s.Size())
	i := s.Size() - 1

	for node := s.head; node != nil; node = node.prev {
		strs[i] = fmt.Sprintf("%v", node.value)
		i--
	}

	sb.WriteString(strings.Join(strs, ","))

	return sb.String()
}

func (s *LinkedStack[T]) Push(value T) {
	s.head = &node[T]{
		value: value,
		prev:  s.head,
	}
	s.size++
}

func (s *LinkedStack[T]) PushAll(values ...T) {
	for _, value := range values {
		s.Push(value)
	}
}

func (s *LinkedStack[T]) Pop() (T, bool) {
	value, ok := s.Peek()
	if ok {
		s.head = s.head.prev
		s.size--
	}

	return value, ok
}

func (s *LinkedStack[T]) Peek() (T, bool) {
	if s.Empty() {
		return *new(T), false
	}

	return s.head.value, true
}
