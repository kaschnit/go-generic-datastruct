package concurrentmap_test

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/concurrentmap"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

func getMapsForTest[K comparable, V any](entries ...entry.Entry[K, V]) []mapp.Map[K, V] {
	return []mapp.Map[K, V]{
		hashmap.New(entries...),
	}
}

func TestConcurrentMapString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		mapping           mapp.Map[string, int]
		expectedFirstLine string
	}{
		{
			name:              "empty hashmap",
			mapping:           hashmap.New[string, int](),
			expectedFirstLine: "HashMap",
		},
		{
			name:              "hashmap with 1 item",
			mapping:           hashmap.New(entry.New("foo", 987654321)),
			expectedFirstLine: "HashMap",
		},
		{
			name: "hashmap with a few items",
			mapping: hashmap.New(
				entry.New("abc", 100),
				entry.New("def", 1145),
				entry.New("ghi", -202),
				entry.New("jkl", 5),
				entry.New("mno", 6),
				entry.New("pqr", 7),
			),
			expectedFirstLine: "HashMap",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			m := concurrentmap.MakeThreadSafe(testCase.mapping)
			resultLines := strings.Split(m.String(), "\n")
			assert.Len(t, resultLines, 2, "expected 2 lines in ConcurrentSet.String() output")
			assert.Equal(t, fmt.Sprintf("[Concurrent]%s", testCase.expectedFirstLine), resultLines[0])

			// Map does not guarantee ordering
			testCase.mapping.ForEach(func(key string, value int) {
				assert.Contains(t, resultLines[1], entry.NewRef(key, value).String())
			})
		})
	}
}

func TestConcurrentMapConcurrentPutAndRemove(t *testing.T) {
	t.Parallel()

	innerMaps := getMapsForTest[string, int]()
	for _, innerMap := range innerMaps {
		t.Run(fmt.Sprintf("%T", innerMap), func(t *testing.T) {
			m := concurrentmap.MakeThreadSafe(innerMap)
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
					m.Put(fmt.Sprintf("%d", value), value)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.Equal(t, size, m.Size())

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(sleepDuration)

					key := fmt.Sprintf("%d", value)

					actualValue, ok := m.Get(key)
					assert.True(t, ok)
					assert.Equal(t, value, actualValue)

					contained := m.RemoveKey(key)
					assert.True(t, contained)
				}()
			}

			// Wait for all goroutines to finish
			wg.Wait()

			assert.True(t, m.Empty())
			assert.Equal(t, 0, m.Size())
		})
	}
}

func TestMakeThreadSafe_AlreadyThreadSafe(t *testing.T) {
	m := hashmap.New[int, string]()
	c1 := concurrentmap.MakeThreadSafe[int, string](m)
	c2 := concurrentmap.MakeThreadSafe[int, string](c1)
	c3 := concurrentmap.MakeThreadSafe[int, string](c2)
	assert.NotEqual(t, m, c1)
	assert.Equal(t, c1, c2)
	assert.Equal(t, c1, c3)
	assert.Equal(t, c2, c3)
}
