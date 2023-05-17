package concurrentlist_test

import (
	"fmt"
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

// Ensure that ConcurrentList implements List.
var _ list.List[int] = &concurrentlist.ConcurrentList[int]{}

func getListsForTest[T any](values ...T) []list.List[T] {
	return []list.List[T]{
		arraylist.New(values...),
		linkedlist.NewSingleLinked(values...),
		linkedlist.NewDoubleLinked(values...),
	}
}

func TestConcurrentListString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		list           list.List[string]
		expectedSuffix string
	}{
		{
			name:           "empty arraylist",
			list:           arraylist.New[string](),
			expectedSuffix: "ArrayList\n",
		},
		{
			name:           "empty single linked list",
			list:           linkedlist.NewSingleLinked[string](),
			expectedSuffix: "SingleLinkedList\n",
		},
		{
			name:           "arraylist with 1 item",
			list:           arraylist.New("foo"),
			expectedSuffix: "ArrayList\nfoo",
		},
		{
			name:           "double linked list with 1 item",
			list:           linkedlist.NewDoubleLinked("foo"),
			expectedSuffix: "DoubleLinkedList\nfoo",
		},
		{
			name:           "arraylist with a few items",
			list:           arraylist.New("abc", "def", "ghi", "jkl", "mno", "pqr"),
			expectedSuffix: "ArrayList\nabc,def,ghi,jkl,mno,pqr",
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			l := concurrentlist.MakeThreadSafe(testCase.list)
			assert.Equal(t, fmt.Sprintf("[Concurrent]%s", testCase.expectedSuffix), l.String())
		})
	}
}

func TestConcurrentListConcurrentAppendAndPop(t *testing.T) {
	t.Parallel()

	innerLists := getListsForTest[int]()
	for i := range innerLists {
		innerList := innerLists[i]
		t.Run(fmt.Sprintf("%T", innerList), func(t *testing.T) {
			t.Parallel()

			l := concurrentlist.MakeThreadSafe(innerList)
			waitGroup := sync.WaitGroup{}

			size := 5000

			for i := 0; i < size; i++ {
				index := i
				value := index * 2

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					time.Sleep(sleepDuration)
					l.Append(value)
				}()
			}

			// Wait for all goroutines to finish
			waitGroup.Wait()

			assert.Equal(t, size, l.Size())

			for i := 0; i < size; i++ {
				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					time.Sleep(sleepDuration)

					actualValue, ok := l.PopBack()
					assert.True(t, ok)
					assert.True(t, actualValue >= 0 && actualValue <= (size-1)*2)
				}()
			}

			// Wait for all goroutines to finish
			waitGroup.Wait()

			assert.True(t, l.Empty())
			assert.Equal(t, 0, l.Size())
		})
	}
}

func TestMakeThreadSafe_AlreadyThreadSafe(t *testing.T) {
	t.Parallel()

	l := arraylist.New[int]()
	c1 := concurrentlist.MakeThreadSafe[int](l)
	c2 := concurrentlist.MakeThreadSafe[int](c1)
	c3 := concurrentlist.MakeThreadSafe[int](c2)

	assert.NotEqual(t, l, c1)
	assert.Equal(t, c1, c2)
	assert.Equal(t, c1, c3)
	assert.Equal(t, c2, c3)
}
