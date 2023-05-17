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
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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
	for i := range innerMaps {
		innerMap := innerMaps[i]
		t.Run(fmt.Sprintf("%T", innerMap), func(t *testing.T) {
			t.Parallel()

			m := concurrentmap.MakeThreadSafe(innerMap)
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
					m.Put(fmt.Sprintf("%d", value), value)
				}()
			}

			// Wait for all goroutines to finish
			waitGroup.Wait()

			assert.Equal(t, size, m.Size())

			for i := 0; i < size; i++ {
				value := i

				// Make the goroutine wait for a random duration between 0.0 and 0.1 seconds
				sleepDuration := time.Duration(rand.Float64()/10) * time.Second

				waitGroup.Add(1)
				go func() {
					defer waitGroup.Done()
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
			waitGroup.Wait()

			assert.True(t, m.Empty())
			assert.Equal(t, 0, m.Size())
		})
	}
}

func TestMakeThreadSafe_AlreadyThreadSafe(t *testing.T) {
	t.Parallel()

	m := hashmap.New[int, string]()
	map1 := concurrentmap.MakeThreadSafe[int, string](m)
	map2 := concurrentmap.MakeThreadSafe[int, string](map1)
	map3 := concurrentmap.MakeThreadSafe[int, string](map2)

	assert.NotEqual(t, m, map1)
	assert.Equal(t, map1, map2)
	assert.Equal(t, map1, map3)
	assert.Equal(t, map2, map3)
}
