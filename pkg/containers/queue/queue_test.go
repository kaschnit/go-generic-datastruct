package queue_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/compare"
	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/linkedqueue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/priorityqueue/heappq"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func getQueuesForTest[T constraints.Ordered](values ...T) []queue.Queue[T] {
	return []queue.Queue[T]{
		heappq.New(compare.OrderedComparator[T], values...),
		linkedqueue.New(values...),
	}
}

func TestEmptyFalse(t *testing.T) {
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			queues := getQueuesForTest(testCase.initial...)
			for _, q := range queues {
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					assert.False(t, q.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	queues := getQueuesForTest[int]()
	for _, q := range queues {
		assert.True(t, q.Empty())

		q.Push(1)
		assert.False(t, q.Empty())

		q.Pop()
		assert.True(t, q.Empty())
	}
}
