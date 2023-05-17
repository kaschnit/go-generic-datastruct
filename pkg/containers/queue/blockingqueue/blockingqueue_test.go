package blockingqueue_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/blockingqueue"
	"github.com/stretchr/testify/assert"
)

var _ queue.Queue[int] = &blockingqueue.BlockingQueue[int]{}

func TestBlockingQueueString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		queue    *blockingqueue.BlockingQueue[any]
		expected string
	}{
		{
			name:     "empty queue",
			queue:    blockingqueue.NewBuilder[any](5).Build(),
			expected: "BlockingQueue[capacity=5]\n",
		},
		{
			name:     "empty queue with no capacity",
			queue:    blockingqueue.NewBuilder[any](0).Build(),
			expected: "BlockingQueue[capacity=0]\n",
		},
		{
			name:     "queue with 1 item",
			queue:    blockingqueue.NewBuilder[any](2).AddItems("987654321").Build(),
			expected: "BlockingQueue[capacity=2]\n987654321",
		},
		{
			name:     "queue with a few items",
			queue:    blockingqueue.NewBuilder[any](6).AddItems(100, 1145, -202, 5, 6, 7).Build(),
			expected: "BlockingQueue[capacity=6]\n7,6,5,-202,1145,100",
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, testCase.expected, testCase.queue.String())
		})
	}
}

func TestEmptyFalse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []string
	}{
		{
			name:    "3 items",
			initial: []string{"a", "b", "c"},
		},
		{
			name:    "1 item",
			initial: []string{"hello"},
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			q := blockingqueue.NewBuilder[string](len(testCase.initial) + 5).
				AddItems(testCase.initial...).
				Build()
			assert.False(t, q.Empty())
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	t.Parallel()

	q := blockingqueue.NewBuilder[int](3).Build()
	assert.True(t, q.Empty())

	q.Push(1)
	assert.False(t, q.Empty())

	q.Pop()
	assert.True(t, q.Empty())
}

func TestSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  []float64
		expected int
	}{
		{
			name:     "3 items",
			initial:  []float64{1.2, 2.3, 999.999},
			expected: 3,
		},
		{
			name:     "1 item",
			initial:  []float64{7.000},
			expected: 1,
		},
		{
			name:     "0 items",
			initial:  []float64{},
			expected: 0,
		},
		{
			name:     "6 items",
			initial:  []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
			expected: 6,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			q := blockingqueue.NewBuilder[float64](len(testCase.initial) + 5).
				AddItems(testCase.initial...).
				Build()
			assert.Equal(t, testCase.expected, q.Size())
			assert.Equal(t, len(testCase.initial), q.Size())
		})
	}
}

func TestClearNonEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []float64
	}{
		{
			name:    "3 items",
			initial: []float64{1.2, 2.3, 999.999},
		},
		{
			name:    "1 item",
			initial: []float64{7.000},
		},
		{
			name:    "6 items",
			initial: []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			q := blockingqueue.NewBuilder[float64](len(testCase.initial) + 5).
				AddItems(testCase.initial...).
				Build()
			assert.False(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())

			q.Push(1.2345)
			assert.False(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())
		})
	}
}

func TestClearEmpty(t *testing.T) {
	t.Parallel()

	q := blockingqueue.NewBuilder[string](3).Build()
	assert.True(t, q.Empty())

	q.Clear()
	assert.True(t, q.Empty())

	q.Clear()
	assert.True(t, q.Empty())

	q.Push("hello")
	assert.False(t, q.Empty())

	q.Clear()
	assert.True(t, q.Empty())
}

func TestPeekOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		items    []int
		expected int
	}{
		{
			name:     "peek at an item",
			items:    []int{9, 8, 7},
			expected: 9,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			q := blockingqueue.NewBuilder[int](len(testCase.items)).
				AddItems(testCase.items...).
				Build()
			item, ok := q.Peek()
			assert.True(t, ok)
			assert.Equal(t, testCase.expected, item)
		})
	}
}

func TestPeekNotOk(t *testing.T) {
	t.Parallel()

	q := blockingqueue.NewBuilder[int](3).Build()

	_, ok := q.Peek()
	assert.False(t, ok)

	q.Push(1)

	item, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, item)

	item, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, item)

	_, ok = q.Peek()
	assert.False(t, ok)
}

func TestAddItems_MoreThanCapacity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "exact same amount",
			items: []int{9, 8, 7},
		},
		{
			name:  "one more than capacity",
			items: []int{9, 8, 7, 6},
		},
		{
			name:  "way more than capacity",
			items: []int{9, 8, 7, 100, 200, 60},
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			q := blockingqueue.NewBuilder[int](3).
				AddItems(testCase.items...).
				Build()

			assert.False(t, q.Empty())
			assert.Equal(t, q.Size(), 3)

			item, ok := q.Pop()
			assert.True(t, ok)
			assert.Equal(t, testCase.items[0], item)

			item, ok = q.Pop()
			assert.True(t, ok)
			assert.Equal(t, testCase.items[1], item)

			item, ok = q.Pop()
			assert.True(t, ok)
			assert.Equal(t, testCase.items[2], item)

			assert.True(t, q.Empty())
			assert.Equal(t, q.Size(), 0)
		})
	}
}
