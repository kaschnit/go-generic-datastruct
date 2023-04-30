package blockingqueue_test

import (
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue/blockingqueue"
	"github.com/stretchr/testify/assert"
)

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
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.queue.String())
		})
	}
}
