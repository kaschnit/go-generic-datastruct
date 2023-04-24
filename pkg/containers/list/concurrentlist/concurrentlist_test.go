package concurrentlist_test

import (
	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/containers/list/concurrentlist"
)

// Ensure that ConcurrentList implements List
var _ list.List[int] = &concurrentlist.ConcurrentList[int]{}
