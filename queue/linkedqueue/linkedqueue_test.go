package linkedqueue_test

import (
	"github.com/kaschnit/go-ds/queue"
	"github.com/kaschnit/go-ds/queue/linkedqueue"
)

// Ensure that LinkedQueue implements Queue
var _ queue.Queue[int] = &linkedqueue.LinkedQueue[int]{}
