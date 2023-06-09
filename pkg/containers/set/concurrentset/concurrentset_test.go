package concurrentset_test

import (
	"fmt"
	"math/rand"
	"strings"
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

func TestConcurrentMapString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		values            set.Set[int]
		expectedFirstLine string
	}{
		{
			name:              "empty hashset",
			values:            hashset.New[int](),
			expectedFirstLine: "HashSet",
		},
		{
			name:              "hashset with 1 item",
			values:            hashset.New(987654321),
			expectedFirstLine: "HashSet",
		},
		{
			name:              "hashset with a few items",
			values:            hashset.New(100, 1145, -202, 5, 6, 7),
			expectedFirstLine: "HashSet",
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			s := concurrentset.MakeThreadSafe(testCase.values)
			resultLines := strings.Split(s.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in ConcurrentSet.String() output")
			assert.Equal(t, fmt.Sprintf("[Concurrent]%s", testCase.expectedFirstLine), resultLines[0])

			// Set does not guarantee ordering
			testCase.values.ForEach(func(_ int, value int) {
				assert.Contains(t, resultLines[1], fmt.Sprintf("%d", value))
			})
		})
	}
}

func TestConcurrentSetConcurrentAddAndRemove(t *testing.T) {
	t.Parallel()

	innerSets := getSetsForTest[int]()
	for i := range innerSets {
		innerSet := innerSets[i]
		t.Run(fmt.Sprintf("%T", innerSet), func(t *testing.T) {
			t.Parallel()

			s := concurrentset.MakeThreadSafe(innerSet)
			waitGroup := sync.WaitGroup{}

			size := 5000

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					time.Sleep(sleepDuration)
					s.Add(value)
				}()
			}

			// Wait for all goroutines to finish
			waitGroup.Wait()

			assert.Equal(t, size, s.Size())

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
					time.Sleep(sleepDuration)

					contained := s.Contains(value)
					assert.True(t, contained)

					removed := s.Remove(value)
					assert.True(t, removed)
				}()
			}

			// Wait for all goroutines to finish
			waitGroup.Wait()

			assert.True(t, s.Empty())
			assert.Equal(t, 0, s.Size())
		})
	}
}

func TestMakeThreadSafe_AlreadyThreadSafe(t *testing.T) {
	t.Parallel()

	s := hashset.New[int]()
	c1 := concurrentset.MakeThreadSafe[int](s)
	c2 := concurrentset.MakeThreadSafe[int](c1)
	c3 := concurrentset.MakeThreadSafe[int](c2)

	assert.NotEqual(t, s, c1)
	assert.Equal(t, c1, c2)
	assert.Equal(t, c1, c3)
	assert.Equal(t, c2, c3)
}
