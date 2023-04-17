package linkedlist

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/enumerable"
	"github.com/kaschnit/go-ds/iterator"
)

type singleLinkedListIterator[T any] struct {
	index int
	node  *singleLinkedNode[T]
}

func (a *singleLinkedListIterator[T]) Key() (key int, ok bool) {
	return a.index, a.node != nil
}

func (a *singleLinkedListIterator[T]) Value() (value T, ok bool) {
	return a.node.value, true
}

func (a *singleLinkedListIterator[T]) Next() (next iterator.ForwardIterator[int, T], ok bool) {
	if !a.HasNext() {
		return nil, false
	}
	return &singleLinkedListIterator[T]{
		index: a.index + 1,
		node:  a.node.next,
	}, true
}

func (a *singleLinkedListIterator[T]) HasNext() bool {
	return a.node.next != nil
}

type singleLinkedNode[T any] struct {
	value T
	next  *singleLinkedNode[T]
}

type SingleLinkedList[T any] struct {
	head *singleLinkedNode[T]
	tail *singleLinkedNode[T]
	size int
}

func NewSingleLinked[T any](values ...T) *SingleLinkedList[T] {
	list := SingleLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
	list.AppendAll(values...)
	return &list
}

func (l *SingleLinkedList[T]) Empty() bool {
	return l.Size() == 0
}

func (l *SingleLinkedList[T]) Size() int {
	return l.size
}

func (l *SingleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *SingleLinkedList[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("SingleLinkedList\n")
	strs := make([]string, l.Size())
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		strs[i] = fmt.Sprintf("%v", node.value)
	}
	sb.WriteString(strings.Join(strs, ","))
	return sb.String()
}

func (l *SingleLinkedList[T]) ForEach(op enumerable.Op[int, T]) {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		op(i, node.value)
	}
}

func (l *SingleLinkedList[T]) Any(predicate enumerable.Predicate[int, T]) bool {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if predicate(i, node.value) {
			return true
		}
	}
	return false
}

func (l *SingleLinkedList[T]) All(predicate enumerable.Predicate[int, T]) bool {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if !predicate(i, node.value) {
			return false
		}
	}
	return true
}

func (l *SingleLinkedList[T]) Find(predicate enumerable.Predicate[int, T]) (key int, value T, ok bool) {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if predicate(i, node.value) {
			return i, node.value, true
		}
	}
	return 0, *new(T), false
}

func (l *SingleLinkedList[T]) Keys(abort <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < l.size; i++ {
			select {
			case ch <- i:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *SingleLinkedList[T]) Values(abort <-chan struct{}) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for node := l.head; node != nil; node = node.next {
			select {
			case ch <- node.value:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *SingleLinkedList[T]) Items(abort <-chan struct{}) <-chan enumerable.KeyValue[int, T] {
	ch := make(chan enumerable.KeyValue[int, T])
	go func() {
		defer close(ch)
		for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
			select {
			case ch <- enumerable.KeyValue[int, T]{
				Key:   i,
				Value: node.value,
			}:
			case <-abort:
				return
			}
		}
	}()
	return ch
}

func (l *SingleLinkedList[T]) Iterator() (iter iterator.ForwardIterator[int, T], ok bool) {
	if l.Empty() {
		return nil, false
	}
	return &singleLinkedListIterator[T]{
		index: 0,
		node:  l.head,
	}, true
}

func (l *SingleLinkedList[T]) Append(value T) {
	newTail := &singleLinkedNode[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newTail
	} else {
		l.tail.next = newTail
	}
	l.tail = newTail
	l.size++
}

func (l *SingleLinkedList[T]) AppendAll(values ...T) {
	for _, value := range values {
		l.Append(value)
	}
}

func (l *SingleLinkedList[T]) Prepend(value T) {
	newHead := &singleLinkedNode[T]{
		value: value,
		next:  l.head,
	}
	if l.tail == nil {
		l.tail = newHead
	}
	l.head = newHead
	l.size++
}

func (l *SingleLinkedList[T]) PrependAll(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		l.Prepend(values[i])
	}
}

func (l *SingleLinkedList[T]) Insert(index int, value T) (ok bool) {
	if index < 0 || index > l.Size() {
		return false
	} else if index == 0 {
		l.Prepend(value)
		return true
	}

	// Find the nodes between which the new node will be placed
	prevNode := l.getNode(index - 1)
	nextNode := prevNode.next

	// Create the new node
	newNode := &singleLinkedNode[T]{
		value: value,
		next:  nextNode,
	}

	// Insert the new node at the insertion point
	prevNode.next = newNode

	l.size++

	return true
}

func (l *SingleLinkedList[T]) InsertAll(index int, values ...T) (ok bool) {
	if len(values) == 0 {
		return true
	} else if index < 0 || index > l.Size() {
		return false
	} else if index == 0 {
		l.PrependAll(values...)
		return true
	}

	// Find the nodes between which the new node will be placed
	prevNode := l.getNode(index - 1)
	nextNode := prevNode.next

	// Create a sub list
	subList := NewSingleLinked(values...)

	// Insert the sub list at the insertion point
	subList.tail.next = nextNode
	prevNode.next = subList.head

	l.size += subList.Size()

	return true
}

func (l *SingleLinkedList[T]) PopBack() (value T, ok bool) {
	back, ok := l.GetBack()
	if ok {
		if l.Size() == 1 {
			l.head = nil
			l.tail = nil
		} else {
			node := &singleLinkedNode[T]{
				next: l.head,
			}
			for i := 0; i < l.size-1; i++ {
				node = node.next
			}
			l.tail = node
			l.tail.next = nil
		}

		l.size--
	}
	return back, ok
}

func (l *SingleLinkedList[T]) PopFront() (value T, ok bool) {
	front, ok := l.GetFront()
	if ok {
		l.head = l.head.next
		l.size--
	}
	return front, ok
}

func (l *SingleLinkedList[T]) GetFront() (value T, ok bool) {
	if l.Empty() {
		return *new(T), false
	}
	return l.head.value, true
}

func (l *SingleLinkedList[T]) GetBack() (value T, ok bool) {
	if l.Empty() {
		return *new(T), false
	}
	return l.tail.value, true
}

func (l *SingleLinkedList[T]) Get(index int) (value T, ok bool) {
	if index < 0 || index >= l.Size() {
		return *new(T), false
	}
	node := l.getNode(index)
	return node.value, true
}

func (l *SingleLinkedList[T]) getNode(index int) *singleLinkedNode[T] {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
