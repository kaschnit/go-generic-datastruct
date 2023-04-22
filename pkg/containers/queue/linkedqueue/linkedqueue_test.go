package linkedqueue_test

import (
	"github.com/kaschnit/go-ds/pkg/containers/queue"
	"github.com/kaschnit/go-ds/pkg/containers/queue/linkedqueue"
)

// Ensure that LinkedQueue implements Queue
var _ queue.Queue[int] = &linkedqueue.LinkedQueue[int]{}
