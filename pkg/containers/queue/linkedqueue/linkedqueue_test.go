package linkedqueue_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/linkedqueue"
	"github.com/stretchr/testify/assert"
)

// Ensure that LinkedQueue implements Queue
var _ queue.Queue[int] = &linkedqueue.LinkedQueue[int]{}

func TestLinkedQueueString(t *testing.T) {
	tests := []struct {
		name     string
		queue    *linkedqueue.LinkedQueue[int]
		expected string
	}{
		{
			name:     "empty queue",
			queue:    linkedqueue.New[int](),
			expected: "LinkedQueue\n",
		},
		{
			name:     "queue with 1 item",
			queue:    linkedqueue.New(987654321),
			expected: "LinkedQueue\n987654321",
		},
		{
			name:     "queue with a few items",
			queue:    linkedqueue.New(100, 1145, -202, 5, 6, 7),
			expected: "LinkedQueue\n7,6,5,-202,1145,100",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.queue.String())
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name         string
		initial      []string
		pushItem     string
		expectedPeek string
	}{
		{
			name:         "3 items pushing a string",
			initial:      []string{"1", "two", "3"},
			pushItem:     "foo",
			expectedPeek: "1",
		},
		{
			name:         "3 items pushing an empty string",
			initial:      []string{"1", "two", "3"},
			pushItem:     "",
			expectedPeek: "1",
		},
		{
			name:         "0 items pushing a string",
			initial:      []string{},
			pushItem:     "hello",
			expectedPeek: "hello",
		},
		{
			name:         "0 items pushing an empty string",
			initial:      []string{},
			pushItem:     "",
			expectedPeek: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			q := linkedqueue.New(testCase.initial...)
			t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
				prevSize := q.Size()
				q.Push(testCase.pushItem)
				assert.Equal(t, prevSize+1, q.Size())

				actual, ok := q.Peek()
				assert.True(t, ok)
				assert.Equal(t, testCase.expectedPeek, actual)
			})
		})
	}
}

func TestPushAll(t *testing.T) {
	tests := []struct {
		name         string
		initial      []int
		pushItems    []int
		expectedPeek int
	}{
		{
			name:         "pushing no items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{},
			expectedPeek: 9,
		},
		{
			name:         "pushing one item",
			initial:      []int{9, 8, 7},
			pushItems:    []int{1},
			expectedPeek: 9,
		},
		{
			name:         "pushing some items",
			initial:      []int{9, 8, 7},
			pushItems:    []int{500, 1000},
			expectedPeek: 9,
		},
		{
			name:         "pushing one item onto an empty queue",
			initial:      []int{},
			pushItems:    []int{-52},
			expectedPeek: -52,
		},
		{
			name:         "pushing some items onto an empty queue",
			initial:      []int{},
			pushItems:    []int{1, 2, 3},
			expectedPeek: 1,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			q := linkedqueue.New(testCase.initial...)
			t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
				prevSize := q.Size()
				q.PushAll(testCase.pushItems...)
				assert.Equal(t, prevSize+len(testCase.pushItems), q.Size())

				actual, ok := q.Peek()
				assert.True(t, ok)
				assert.Equal(t, testCase.expectedPeek, actual)
			})
		})
	}
}
