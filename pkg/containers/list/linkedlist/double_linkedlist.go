package linkedlist

import (
	"fmt"
	"strings"

	"github.com/kaschnit/go-ds/pkg/containers/enumerable"
	"github.com/kaschnit/go-ds/pkg/iterator"
)

type doubleLinkedListIterator[T any] struct {
	index  int
	node   *doubleLinkedNode[T]
	nextOp func(index int, node *doubleLinkedNode[T]) (int, *doubleLinkedNode[T])
}

func (a *doubleLinkedListIterator[T]) Key() (int, bool) {
	return a.index, a.node != nil
}

func (a *doubleLinkedListIterator[T]) Value() (T, bool) {
	return a.node.value, true
}

func (a *doubleLinkedListIterator[T]) Next() (iterator.ForwardIterator[int, T], bool) {
	if !a.HasNext() {
		return nil, false
	}

	nextIndex, nextNode := a.nextOp(a.index, a.node)

	return &doubleLinkedListIterator[T]{
		index:  nextIndex,
		node:   nextNode,
		nextOp: a.nextOp,
	}, true
}

func (a *doubleLinkedListIterator[T]) HasNext() bool {
	_, nextNode := a.nextOp(a.index, a.node)

	return nextNode != nil
}

type doubleLinkedNode[T any] struct {
	value T
	prev  *doubleLinkedNode[T]
	next  *doubleLinkedNode[T]
}

type DoubleLinkedList[T any] struct {
	head *doubleLinkedNode[T]
	tail *doubleLinkedNode[T]
	size int
}

func NewDoubleLinked[T any](values ...T) *DoubleLinkedList[T] {
	list := DoubleLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
	list.AppendAll(values...)

	return &list
}

func (l *DoubleLinkedList[T]) Empty() bool {
	return l.Size() == 0
}

func (l *DoubleLinkedList[T]) Size() int {
	return l.size
}

func (l *DoubleLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *DoubleLinkedList[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("DoubleLinkedList\n")

	strs := make([]string, l.Size())
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		strs[i] = fmt.Sprintf("%v", node.value)
	}

	sb.WriteString(strings.Join(strs, ","))

	return sb.String()
}

func (l *DoubleLinkedList[T]) ForEach(op enumerable.Op[int, T]) {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		op(i, node.value)
	}
}

func (l *DoubleLinkedList[T]) Any(predicate enumerable.Predicate[int, T]) bool {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if predicate(i, node.value) {
			return true
		}
	}

	return false
}

func (l *DoubleLinkedList[T]) All(predicate enumerable.Predicate[int, T]) bool {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if !predicate(i, node.value) {
			return false
		}
	}

	return true
}

func (l *DoubleLinkedList[T]) Find(predicate enumerable.Predicate[int, T]) (int, T, bool) {
	for i, node := 0, l.head; node != nil; i, node = i+1, node.next {
		if predicate(i, node.value) {
			return i, node.value, true
		}
	}

	return 0, *new(T), false
}

func (l *DoubleLinkedList[T]) Iterator() (iterator.ForwardIterator[int, T], bool) {
	if l.Empty() {
		return nil, false
	}

	return &doubleLinkedListIterator[T]{
		index: 0,
		node:  l.head,
		nextOp: func(index int, node *doubleLinkedNode[T]) (int, *doubleLinkedNode[T]) {
			return index + 1, node.next
		},
	}, true
}

func (l *DoubleLinkedList[T]) IteratorReverse() (iterator.ForwardIterator[int, T], bool) {
	if l.Empty() {
		return nil, false
	}

	return &doubleLinkedListIterator[T]{
		index: l.Size() - 1,
		node:  l.tail,
		nextOp: func(index int, node *doubleLinkedNode[T]) (int, *doubleLinkedNode[T]) {
			return index - 1, node.prev
		},
	}, true
}

func (l *DoubleLinkedList[T]) Append(value T) {
	newTail := &doubleLinkedNode[T]{
		value: value,
		prev:  l.tail,
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

func (l *DoubleLinkedList[T]) AppendAll(values ...T) {
	for _, value := range values {
		l.Append(value)
	}
}

func (l *DoubleLinkedList[T]) Prepend(value T) {
	newHead := &doubleLinkedNode[T]{
		value: value,
		prev:  nil,
		next:  l.head,
	}
	if l.tail == nil {
		l.tail = newHead
	} else {
		l.head.prev = newHead
	}

	l.head = newHead
	l.size++
}

func (l *DoubleLinkedList[T]) PrependAll(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		l.Prepend(values[i])
	}
}

func (l *DoubleLinkedList[T]) Insert(index int, value T) bool {
	if index < 0 || index > l.Size() {
		return false
	} else if index == 0 {
		l.Prepend(value)

		return true
	} else if index == l.Size() {
		l.Append(value)

		return true
	}

	// Find the nodes between which the new node will be placed
	prevNode := l.getNode(index - 1)
	nextNode := prevNode.next

	// Create the new node
	newNode := &doubleLinkedNode[T]{
		value: value,
		prev:  prevNode,
		next:  nextNode,
	}

	// Insert the new node at the insertion point.
	// prevNode's index was validated, so it will not be nil.
	// nextNode may be nil, since prevNode is allowed to be the last node in the list.
	prevNode.next = newNode

	if nextNode != nil {
		nextNode.prev = newNode
	}

	l.size++

	return true
}

func (l *DoubleLinkedList[T]) InsertAll(index int, values ...T) bool {
	if len(values) == 0 {
		return true
	} else if index < 0 || index > l.Size() {
		return false
	} else if index == 0 {
		l.PrependAll(values...)

		return true
	} else if index == l.Size() {
		l.AppendAll(values...)

		return true
	}

	// Find the nodes between which the new node will be placed
	prevNode := l.getNode(index - 1)
	nextNode := prevNode.next

	// Create a sub list
	subList := NewDoubleLinked(values...)

	// Insert the sub list at the insertion point.
	// prevNode's index was validated, so it will not be nil.
	// subList.head and subList.tail will not be nil because subList was checked to be non-empty.
	// nextNode may be nil, since prevNode is allowed to be the last node in the list.
	prevNode.next = subList.head
	subList.head.prev = prevNode

	if nextNode != nil {
		subList.tail.next = nextNode
		nextNode.prev = subList.tail
	}

	l.size += subList.Size()

	return true
}

func (l *DoubleLinkedList[T]) PopBack() (T, bool) {
	back, ok := l.GetBack()
	if ok {
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		} else {
			l.tail.next = nil
		}

		l.size--
	}

	return back, ok
}

func (l *DoubleLinkedList[T]) PopFront() (T, bool) {
	front, ok := l.GetFront()
	if ok {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		} else {
			l.head.prev = nil
		}
		l.size--
	}

	return front, ok
}

func (l *DoubleLinkedList[T]) GetFront() (T, bool) {
	if l.Empty() {
		return *new(T), false
	}

	return l.head.value, true
}

func (l *DoubleLinkedList[T]) GetBack() (T, bool) {
	if l.Empty() {
		return *new(T), false
	}

	return l.tail.value, true
}

func (l *DoubleLinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.Size() {
		return *new(T), false
	}

	return l.getNode(index).value, true
}

func (l *DoubleLinkedList[T]) getNode(index int) *doubleLinkedNode[T] {
	if index < l.Size()/2 {
		return l.getNodeFromFront(index)
	}

	return l.getNodeFromBack(index)
}

func (l *DoubleLinkedList[T]) getNodeFromFront(index int) *doubleLinkedNode[T] {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *DoubleLinkedList[T]) getNodeFromBack(index int) *doubleLinkedNode[T] {
	node := l.tail
	for i := l.Size() - 1; i > index; i-- {
		node = node.prev
	}

	return node
}
