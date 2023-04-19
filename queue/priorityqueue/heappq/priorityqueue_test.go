package heappq_test

import (
	"github.com/kaschnit/go-ds/queue"
	"github.com/kaschnit/go-ds/queue/priorityqueue/heappq"
)

// Ensure that HeapPQ implements Queue
var _ queue.Queue[int] = &heappq.HeapPQ[int]{}
