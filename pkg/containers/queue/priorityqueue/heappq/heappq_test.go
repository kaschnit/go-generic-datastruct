package heappq_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/priorityqueue/heappq"
	"github.com/stretchr/testify/assert"
)

// Ensure that HeapPQ implements Queue
var _ queue.Queue[int] = &heappq.HeapPQ[int]{}

func TestHeapPQString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		queue    *heappq.HeapPQ[int]
		expected string
	}{
		{
			name:     "empty queue",
			queue:    heappq.New[int](),
			expected: "HeapPQ\n",
		},
		{
			name:     "queue with 1 item",
			queue:    heappq.New(987654321),
			expected: "HeapPQ\n987654321",
		},
		{
			name:     "queue with a few items",
			queue:    heappq.New(100, 1145, -202, 5, 6, 7),
			expected: "HeapPQ\n1145,100,7,6,5,-202",
		},
		{
			name:     "queue with ascending items",
			queue:    heappq.New(1, 2, 3, 4, 5),
			expected: "HeapPQ\n5,4,3,2,1",
		},
		{
			name:     "queue with more ascending items",
			queue:    heappq.New(1, 2, 3, 4, 5, 6, 7, 8),
			expected: "HeapPQ\n8,7,6,5,4,3,2,1",
		},
		{
			name:     "queue with descending items",
			queue:    heappq.New(10, 8, 6, 4, 2),
			expected: "HeapPQ\n10,8,6,4,2",
		},
		{
			name:     "queue with more descending items",
			queue:    heappq.New(10, 8, 6, 4, 2, 0, -2, -4),
			expected: "HeapPQ\n10,8,6,4,2,0,-2,-4",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Assert twice because HeapPQ.String() makes a copy of the heap, want
			// to make sure the copy is not sharing data.
			assert.Equal(t, testCase.expected, testCase.queue.String())
			assert.Equal(t, testCase.expected, testCase.queue.String())
		})
	}
}

func TestHeapPQKitchenSink(t *testing.T) {
	t.Parallel()

	// []
	q := heappq.New[int]()
	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Size())

	// Before []
	// After []
	for i := 0; i < 5; i++ {
		_, ok := q.Peek()
		assert.False(t, ok)

		_, ok = q.Pop()
		assert.False(t, ok)

		_, ok = q.Peek()
		assert.False(t, ok)
	}
	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Size())

	// Before []
	// After [1]
	q.Push(1)
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Size())
	for i := 0; i < 3; i++ {
		value, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, 1, value)
	}

	// Before [1]
	// After [30, 20, 10, 1]
	q.Push(20)
	q.Push(10)
	q.Push(30)
	assert.False(t, q.Empty())
	assert.Equal(t, 4, q.Size())

	for i := 0; i < 3; i++ {
		value, ok := q.Peek()
		assert.True(t, ok)
		assert.Equal(t, 30, value)
	}

	// Before [30, 20, 10, 1]
	// After [20, 10, 1]
	value, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 30, value)

	// Before [20, 10, 1]
	// After [10, 1]
	value, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 20, value)

	// Before [10, 1]
	// After [1]
	value, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 10, value)

	// Before [1]
	// After []
	value, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	// Before [1]
	// After []
	_, ok = q.Pop()
	assert.False(t, ok)

	// Before []
	// After [200, 100, 77, 6, 3, 2, 2, 1]
	q.PushAll(1, 2, 200, 100, 3, 6, 77, 2)

	// Before []
	// After [6, 3, 2, 2, 1]
	q.Pop()
	q.Pop()
	q.Pop()

	// Before []
	// After [6, 5, 3, 2, 2, 1, -10]
	q.Push(-10)
	q.Push(5)

	value, _ = q.Pop()
	assert.Equal(t, 6, value)

	value, _ = q.Pop()
	assert.Equal(t, 5, value)

	value, _ = q.Pop()
	assert.Equal(t, 3, value)

	value, _ = q.Pop()
	assert.Equal(t, 2, value)

	value, _ = q.Pop()
	assert.Equal(t, 2, value)

	value, _ = q.Pop()
	assert.Equal(t, 1, value)

	value, _ = q.Pop()
	assert.Equal(t, -10, value)
}
