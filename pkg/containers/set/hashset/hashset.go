package hashset

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
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

func (s *HashSet[T]) ForEach(op enumerable.Op[T, T]) {
	for value := range s.values {
		op(value, value)
	}
}

func (s *HashSet[T]) Any(predicate enumerable.Predicate[T, T]) bool {
	for value := range s.values {
		if predicate(value, value) {
			return true
		}
	}
	return false
}

func (s *HashSet[T]) All(predicate enumerable.Predicate[T, T]) bool {
	for value := range s.values {
		if !predicate(value, value) {
			return false
		}
	}
	return true
}

func (s *HashSet[T]) Find(predicate enumerable.Predicate[T, T]) (key T, value T, ok bool) {
	for value := range s.values {
		if predicate(value, value) {
			return value, value, true
		}
	}
	return *new(T), *new(T), false
}

func (s *HashSet[T]) Keys(abort <-chan struct{}) <-chan T {
	return s.Values(abort)
}

func (s *HashSet[T]) Values(abort <-chan struct{}) <-chan T {
	ch := make(chan T, 1)
	go func() {
		defer close(ch)
		for value := range s.values {
			select {
			case ch <- value:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (s *HashSet[T]) Items(abort <-chan struct{}) <-chan enumerable.KeyValue[T, T] {
	ch := make(chan enumerable.KeyValue[T, T], 1)
	go func() {
		defer close(ch)
		for value := range s.values {
			select {
			case ch <- enumerable.KeyValue[T, T]{
				Key:   value,
				Value: value,
			}:
			case <-abort:
				return
			}
		}
	}()
	return ch
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

func (s *HashSet[T]) RemoveAll(values ...T) int {
	removed := 0
	for _, value := range values {
		if s.Remove(value) {
			removed++
		}
	}
	return removed
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
