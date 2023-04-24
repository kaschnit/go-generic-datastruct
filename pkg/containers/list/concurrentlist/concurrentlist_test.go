package concurrentlist_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/kaschnit/go-ds/pkg/containers/list"
	"github.com/kaschnit/go-ds/pkg/containers/list/arraylist"
	"github.com/kaschnit/go-ds/pkg/containers/list/concurrentlist"
	"github.com/kaschnit/go-ds/pkg/containers/list/linkedlist"
	"github.com/stretchr/testify/assert"
)

// Ensure that ConcurrentList implements List
var _ list.List[int] = &concurrentlist.ConcurrentList[int]{}

func getListsForTest[T any](values ...T) []list.List[T] {
	return []list.List[T]{
		arraylist.New(values...),
		linkedlist.NewSingleLinked(values...),
		linkedlist.NewDoubleLinked(values...),
	}
}

func TestConcurrentListConcurrentAppendAndPop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []int
	}{
		{
			name:  "empty list",
			items: []int{},
		},
	}

	for _, testCase := range tests {
		innerLists := getListsForTest(testCase.items...)
		for _, innerList := range innerLists {
			l := concurrentlist.MakeThreadSafe(innerList)
			wg := sync.WaitGroup{}

			size := 5000

			for i := 0; i < size; i++ {
				index := i
				value := index * 2

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(sleepDuration)
					l.Append(value)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.Equal(t, size, l.Size())

			for i := 0; i < size; i++ {
				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(sleepDuration)

					actualValue, ok := l.PopBack()
					assert.True(t, ok)
					assert.True(t, actualValue >= 0 && actualValue <= (size-1)*2)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.True(t, l.Empty())
			assert.Equal(t, 0, l.Size())
		}
	}
}
