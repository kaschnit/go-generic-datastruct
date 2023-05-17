package queue_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/concurrentqueue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/linkedqueue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/priorityqueue/heappq"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func getQueuesForTest[T constraints.Ordered](values ...T) []queue.Queue[T] {
	return []queue.Queue[T]{
		heappq.New(values...),
		linkedqueue.New(values...),
		concurrentqueue.MakeThreadSafe[T](linkedqueue.New(values...)),
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

			queues := getQueuesForTest(testCase.initial...)
			for i := range queues {
				q := queues[i]
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					t.Parallel()

					assert.False(t, q.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	t.Parallel()

	queues := getQueuesForTest[int]()
	for i := range queues {
		q := queues[i]
		t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
			t.Parallel()

			assert.True(t, q.Empty())

			q.Push(1)
			assert.False(t, q.Empty())

			q.Pop()
			assert.True(t, q.Empty())
		})
	}
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

			queues := getQueuesForTest(testCase.initial...)
			for i := range queues {
				q := queues[i]
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					t.Parallel()

					assert.Equal(t, testCase.expected, q.Size())
					assert.Equal(t, len(testCase.initial), q.Size())
				})
			}
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

			queues := getQueuesForTest(testCase.initial...)
			for i := range queues {
				q := queues[i]
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					t.Parallel()

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
		})
	}
}

func TestClearEmpty(t *testing.T) {
	t.Parallel()

	queues := getQueuesForTest[string]()
	for i := range queues {
		q := queues[i]
		t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
			t.Parallel()

			assert.True(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())

			q.Push("hello")
			assert.False(t, q.Empty())

			q.Clear()
			assert.True(t, q.Empty())
		})
	}
}
