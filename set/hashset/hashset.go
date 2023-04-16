package hashset

import (
	"fmt"
	"strings"
)

var itemExists = struct{}{}

type HashSet[T comparable] struct {
	values map[T]struct{}
}

func New[T comparable](values ...T) *HashSet[T] {
	set := HashSet[T]{
		values: make(map[T]struct{}),
	}
	set.AddAll(values...)
	return &set
}

func (s *HashSet[T]) Empty() bool {
	return s.Size() == 0
}

func (s *HashSet[T]) Size() int {
	return len(s.values)
}

func (s *HashSet[T]) Clear() {
	for k := range s.values {
		delete(s.values, k)
	}
}

func (s *HashSet[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("HashSet\n")
	strs := []string{}
	for k := range s.values {
		strs = append(strs, fmt.Sprintf("%v", k))
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (s *HashSet[T]) Add(value T) {
	s.values[value] = itemExists
}

func (s *HashSet[T]) AddAll(values ...T) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *HashSet[T]) Contains(value T) bool {
	_, contains := s.values[value]
	return contains
}

func (s *HashSet[T]) Remove(value T) bool {
	contained := s.Contains(value)
	delete(s.values, value)
	return contained
}

func (s *HashSet[T]) RemoveAll(values ...T) bool {
	contained := true
	for _, value := range values {
		contained = contained && s.Remove(value)
	}
	return contained
}

func (s *HashSet[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) ContainsAny(values ...T) bool {
	for _, value := range values {
		if s.Contains(value) {
			return true
		}
	}
	return false
}
