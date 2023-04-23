package heappq_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/compare"
	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/priorityqueue/heappq"
	"github.com/stretchr/testify/assert"
)

// Ensure that HeapPQ implements Queue
var _ queue.Queue[int] = &heappq.HeapPQ[int]{}

func TestHeapPQString(t *testing.T) {
	tests := []struct {
		name     string
		queue    *heappq.HeapPQ[int]
		expected string
	}{
		{
			name:     "empty queue",
			queue:    heappq.New(compare.OrderedComparator[int]),
			expected: "HeapPQ\n",
		},
		{
			name:     "queue with 1 item",
			queue:    heappq.New(compare.OrderedComparator[int], 987654321),
			expected: "HeapPQ\n987654321",
		},
		{
			name:     "queue with a few items",
			queue:    heappq.New(compare.OrderedComparator[int], 100, 1145, -202, 5, 6, 7),
			expected: "HeapPQ\n1145,100,7,6,5,-202",
		},
		{
			name:     "queue with ascending items",
			queue:    heappq.New(compare.OrderedComparator[int], 1, 2, 3, 4, 5),
			expected: "HeapPQ\n5,4,3,2,1",
		},
		{
			name:     "queue with more ascending items",
			queue:    heappq.New(compare.OrderedComparator[int], 1, 2, 3, 4, 5, 6, 7, 8),
			expected: "HeapPQ\n8,7,6,5,4,3,2,1",
		},
		{
			name:     "queue with descending items",
			queue:    heappq.New(compare.OrderedComparator[int], 10, 8, 6, 4, 2),
			expected: "HeapPQ\n10,8,6,4,2",
		},
		{
			name:     "queue with more descending items",
			queue:    heappq.New(compare.OrderedComparator[int], 10, 8, 6, 4, 2, 0, -2, -4),
			expected: "HeapPQ\n10,8,6,4,2,0,-2,-4",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.queue.String())
		})
	}
}
