package concurrentqueue_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/concurrentqueue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/linkedqueue"
	"github.com/stretchr/testify/assert"
)

var _ queue.Queue[int] = &concurrentqueue.ConcurrentQueue[int]{}

func TestConcurrentListString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		queue          queue.Queue[string]
		expectedSuffix string
	}{
		{
			name:           "empty arraylist",
			queue:          linkedqueue.New[string](),
			expectedSuffix: "LinkedQueue\n",
		},
		{
			name:           "arraylist with 1 item",
			queue:          linkedqueue.New("foo"),
			expectedSuffix: "LinkedQueue\nfoo",
		},
		{
			name:           "arraylist with a few items",
			queue:          linkedqueue.New("abc", "def", "ghi", "jkl", "mno", "pqr"),
			expectedSuffix: "LinkedQueue\npqr,mno,jkl,ghi,def,abc",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			l := concurrentqueue.MakeThreadSafe(testCase.queue)
			assert.Equal(t, fmt.Sprintf("[Concurrent]%s", testCase.expectedSuffix), l.String())
		})
	}
}
