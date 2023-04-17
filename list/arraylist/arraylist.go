package arraylist

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/enumerable"
	"github.com/kaschnit/go-ds/iterator"
)

type arrayListIterator[T any] struct {
	index  int
	arr    *ArrayList[T]
	nextOp func(index int) int
}

func (a *arrayListIterator[T]) Key() (key int, ok bool) {
	return a.index, a.index >= 0 && a.index < a.arr.Size()
}

func (a *arrayListIterator[T]) Value() (value T, ok bool) {
	return a.arr.Get(a.index)
}

func (a *arrayListIterator[T]) Next() (next iterator.ForwardIterator[int, T], ok bool) {
	if !a.HasNext() {
		return nil, false
	}
	return &arrayListIterator[T]{
		index:  a.nextOp(a.index),
		arr:    a.arr,
		nextOp: a.nextOp,
	}, true
}

func (a *arrayListIterator[T]) HasNext() bool {
	nextIndex := a.nextOp(a.index)
	return nextIndex >= 0 && nextIndex < a.arr.Size()
}

type ArrayList[T any] struct {
	values []T
}

func New[T any](values ...T) *ArrayList[T] {
	list := ArrayList[T]{
		values: make([]T, len(values)),
	}
	copy(list.values, values)
	return &list
}

func (l *ArrayList[T]) Empty() bool {
	return l.Size() == 0
}

func (l *ArrayList[T]) Size() int {
	return len(l.values)
}

func (l *ArrayList[T]) Clear() {
	l.values = []T{}
}

func (l *ArrayList[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("ArrayList\n")
	strs := []string{}
	for _, value := range l.values {
		strs = append(strs, fmt.Sprintf("%v", value))
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (l *ArrayList[T]) ForEach(op enumerable.Op[int, T]) {
	for i, value := range l.values {
		op(i, value)
	}
}

func (l *ArrayList[T]) Any(predicate enumerable.Predicate[int, T]) bool {
	for i, value := range l.values {
		if predicate(i, value) {
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) All(predicate enumerable.Predicate[int, T]) bool {
	for i, value := range l.values {
		if !predicate(i, value) {
			return false
		}
	}
	return true
}

func (l *ArrayList[T]) Find(predicate enumerable.Predicate[int, T]) (key int, value T, ok bool) {
	for i, value := range l.values {
		if predicate(i, value) {
			return i, value, true
		}
	}
	return 0, *new(T), false
}

func (l *ArrayList[T]) Keys(abort <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range l.values {
			select {
			case ch <- i:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *ArrayList[T]) Values(abort <-chan struct{}) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, value := range l.values {
			select {
			case ch <- value:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *ArrayList[T]) Items(abort <-chan struct{}) <-chan enumerable.KeyValue[int, T] {
	ch := make(chan enumerable.KeyValue[int, T])
	go func() {
		defer close(ch)
		for i, value := range l.values {
			select {
			case ch <- enumerable.KeyValue[int, T]{
				Key:   i,
				Value: value,
			}:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *ArrayList[T]) Iterator() (iter iterator.ForwardIterator[int, T], ok bool) {
	if l.Empty() {
		return nil, false
	}
	return &arrayListIterator[T]{
		index: 0,
		arr:   l,
		nextOp: func(index int) int {
			return index + 1
		},
	}, true
}

func (l *ArrayList[T]) IteratorReverse() (iter iterator.ForwardIterator[int, T], ok bool) {
	if l.Empty() {
		return nil, false
	}
	return &arrayListIterator[T]{
		index: len(l.values) - 1,
		arr:   l,
		nextOp: func(index int) int {
			return index - 1
		},
	}, true
}

func (l *ArrayList[T]) Append(value T) {
	l.values = append(l.values, value)
}

func (l *ArrayList[T]) AppendAll(values ...T) {
	l.values = append(l.values, values...)
}

func (l *ArrayList[T]) Prepend(value T) {
	l.values = append([]T{value}, l.values...)
}

func (l *ArrayList[T]) PrependAll(values ...T) {
	l.values = append(values, l.values...)
}

func (l *ArrayList[T]) Insert(index int, value T) (ok bool) {
	if index < 0 || index > len(l.values) {
		return false
	}
	l.values = append(l.values[:index], append([]T{value}, l.values[index:]...)...)
	return true
}

func (l *ArrayList[T]) InsertAll(index int, values ...T) (ok bool) {
	if index < 0 || index > len(l.values) {
		return false
	}
	l.values = append(l.values[:index], append(values, l.values[index:]...)...)
	return true
}

func (l *ArrayList[T]) PopBack() (value T, ok bool) {
	back, ok := l.GetBack()
	if ok {
		l.values = l.values[:len(l.values)-1]
	}
	return back, ok
}

func (l *ArrayList[T]) PopFront() (value T, ok bool) {
	front, ok := l.GetFront()
	if ok {
		l.values = l.values[1:]
	}
	return front, ok
}

func (l *ArrayList[T]) GetFront() (value T, ok bool) {
	if l.Empty() {
		return *new(T), false
	}
	return l.values[0], true
}

func (l *ArrayList[T]) GetBack() (value T, ok bool) {
	if l.Empty() {
		return *new(T), false
	}
	return l.values[len(l.values)-1], true
}

func (l *ArrayList[T]) Get(index int) (value T, ok bool) {
	if index < 0 || index >= len(l.values) {
		return *new(T), false
	}
	return l.values[index], true
}
