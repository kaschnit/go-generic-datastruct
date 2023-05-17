package arraylist

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/containers/slice"
	"github.com/kaschnit/go-ds/pkg/iterator"
)

type arrayListIterator[T any] struct {
	index  int
	arr    *ArrayList[T]
	nextOp func(index int) int
}

func (a *arrayListIterator[T]) Key() (int, bool) {
	return a.index, a.index >= 0 && a.index < a.arr.Size()
}

func (a *arrayListIterator[T]) Value() (T, bool) {
	return a.arr.Get(a.index)
}

func (a *arrayListIterator[T]) Next() (iterator.ForwardIterator[int, T], bool) {
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
	values slice.Slice[T]
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
	l.values = make(slice.Slice[T], 0)
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
	l.values.ForEach(op)
}

func (l *ArrayList[T]) Any(predicate enumerable.Predicate[int, T]) bool {
	return l.values.Any(predicate)
}

func (l *ArrayList[T]) All(predicate enumerable.Predicate[int, T]) bool {
	return l.values.All(predicate)
}

func (l *ArrayList[T]) Find(predicate enumerable.Predicate[int, T]) (int, T, bool) {
	return l.values.Find(predicate)
}

func (l *ArrayList[T]) Iterator() (iterator.ForwardIterator[int, T], bool) {
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

func (l *ArrayList[T]) IteratorReverse() (iterator.ForwardIterator[int, T], bool) {
	if l.Empty() {
		return nil, false
	}

	return &arrayListIterator[T]{
		index: l.Size() - 1,
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

func (l *ArrayList[T]) Insert(index int, value T) bool {
	if index < 0 || index > l.Size() {
		return false
	}

	l.values = append(l.values[:index], append([]T{value}, l.values[index:]...)...)

	return true
}

func (l *ArrayList[T]) InsertAll(index int, values ...T) bool {
	if index < 0 || index > l.Size() {
		return false
	}

	l.values = append(l.values[:index], append(values, l.values[index:]...)...)

	return true
}

func (l *ArrayList[T]) PopBack() (T, bool) {
	back, ok := l.GetBack()
	if ok {
		l.values = l.values[:len(l.values)-1]
	}

	return back, ok
}

func (l *ArrayList[T]) PopFront() (T, bool) {
	front, ok := l.GetFront()
	if ok {
		l.values = l.values[1:]
	}

	return front, ok
}

func (l *ArrayList[T]) GetFront() (T, bool) {
	if l.Empty() {
		return *new(T), false
	}

	return l.values[0], true
}

func (l *ArrayList[T]) GetBack() (T, bool) {
	if l.Empty() {
		return *new(T), false
	}

	return l.values[len(l.values)-1], true
}

func (l *ArrayList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.Size() {
		return *new(T), false
	}

	return l.values[index], true
}
