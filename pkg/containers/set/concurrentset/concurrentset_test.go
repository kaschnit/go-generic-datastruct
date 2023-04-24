package concurrentset_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/kaschnit/go-ds/pkg/containers/set"
	"github.com/kaschnit/go-ds/pkg/containers/set/concurrentset"
	"github.com/kaschnit/go-ds/pkg/containers/set/hashset"
	"github.com/stretchr/testify/assert"
)

var _ set.Set[int] = &concurrentset.ConcurrentSet[int]{}

func getSetsForTest[T comparable](values ...T) []set.Set[T] {
	return []set.Set[T]{
		hashset.New(values...),
	}
}

func TestConcurrentSetConcurrentAddAndRemove(t *testing.T) {
	t.Parallel()

	innerSets := getSetsForTest[int]()
	for _, innerSet := range innerSets {
		t.Run(fmt.Sprintf("%T", innerSet), func(t *testing.T) {
			s := concurrentset.MakeThreadSafe(innerSet)
			wg := sync.WaitGroup{}

			size := 5000

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(sleepDuration)
					s.Add(value)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.Equal(t, size, s.Size())

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(sleepDuration)

					contained := s.Contains(value)
					assert.True(t, contained)

					removed := s.Remove(value)
					assert.True(t, removed)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.True(t, s.Empty())
			assert.Equal(t, 0, s.Size())
		})
	}

}
