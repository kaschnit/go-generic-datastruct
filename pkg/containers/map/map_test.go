package mapp_test

import (
	"fmt"
	"testing"

	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/concurrentmap"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

func getMapsForTest[K comparable, V any](entries ...entry.Entry[K, V]) []mapp.Map[K, V] {
	return []mapp.Map[K, V]{
		hashmap.New(entries...),
		concurrentmap.MakeThreadSafe[K, V](hashmap.New(entries...)),
	}
}

func TestEmptyFalse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []entry.Entry[string, int]
	}{
		{
			name: "3 items",
			initial: []entry.Entry[string, int]{
				entry.New("a", 1),
				entry.New("b", 2),
				entry.New("c", 3),
			},
		},
		{
			name:    "1 item",
			initial: []entry.Entry[string, int]{entry.New("hello", 123456)},
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					assert.False(t, m.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	t.Parallel()

	maps := getMapsForTest[string, int]()
	for i := range maps {
		m := maps[i]
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			t.Parallel()

			assert.True(t, m.Empty())

			m.Put("foo", 100)
			assert.False(t, m.Empty())

			m.RemoveKey("foo")
			assert.True(t, m.Empty())
		})
	}
}

func TestSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  []entry.Entry[string, float64]
		expected int
	}{
		{
			name: "3 items",
			initial: []entry.Entry[string, float64]{
				entry.New("val1", 1.2),
				entry.New("something", 2.3),
				entry.New("other", 999.999),
			},
			expected: 3,
		},
		{
			name: "3 items with duplicate item added",
			initial: []entry.Entry[string, float64]{
				entry.New("val1", 1.2),
				entry.New("val1", 2.3),
				entry.New("other", 999.999),
			},
			expected: 2,
		},
		{
			name:     "1 item",
			initial:  []entry.Entry[string, float64]{entry.New("foo", 7.000)},
			expected: 1,
		},
		{
			name:     "0 items",
			initial:  []entry.Entry[string, float64]{},
			expected: 0,
		},
		{
			name: "6 items",
			initial: []entry.Entry[string, float64]{
				entry.New("abcdefg", 2.5),
				entry.New("a", 1.000),
				entry.New("no", -5.444),
				entry.New("zyxwvut", 0.1),
				entry.New[string, float64]("", 500),
				entry.New[string, float64]("b", 12),
			},
			expected: 6,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					assert.Equal(t, testCase.expected, m.Size())
				})
			}
		})
	}
}

func TestClearNonEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []entry.Entry[string, float64]
	}{
		{
			name: "3 items",
			initial: []entry.Entry[string, float64]{
				entry.New("eeeee", 1.2),
				entry.New("19", 2.3),
				entry.New("999", 999.999),
			},
		},
		{
			name:    "1 item",
			initial: []entry.Entry[string, float64]{entry.New("ok", 79.999999)},
		},
		{
			name: "6 items",
			initial: []entry.Entry[string, float64]{
				entry.New("y", 2.5),
				entry.New("f", 1.000),
				entry.New("ff", -5.444),
				entry.New("zy", 0.1),
				entry.New("ab", 500.0),
				entry.New("qwerty", 12.0000),
			},
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					assert.False(t, m.Empty())

					m.Clear()
					assert.True(t, m.Empty())

					m.Clear()
					assert.True(t, m.Empty())

					m.Put("0827543982503", 1.2345)
					assert.False(t, m.Empty())

					m.Clear()
					assert.True(t, m.Empty())
				})
			}
		})
	}
}

func TestClearEmpty(t *testing.T) {
	t.Parallel()

	maps := getMapsForTest[string, int]()
	for i := range maps {
		m := maps[i]
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			t.Parallel()

			assert.True(t, m.Empty())

			m.Clear()
			assert.True(t, m.Empty())

			m.Clear()
			assert.True(t, m.Empty())

			m.Put("hello", 3)
			assert.False(t, m.Empty())

			m.Clear()
			assert.True(t, m.Empty())

			m.Put("", 0)
			assert.False(t, m.Empty())

			m.Clear()
			assert.True(t, m.Empty())
		})
	}
}

func TestPutNewKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []entry.Entry[string, int]
		putItem entry.Entry[string, int]
	}{
		{
			name: "3 items adding a string key",
			initial: []entry.Entry[string, int]{
				entry.New("1", 1),
				entry.New("two", 9),
				entry.New("3", 0),
			},
			putItem: entry.New("foo", 3),
		},
		{
			name: "3 items adding an empty string key",
			initial: []entry.Entry[string, int]{
				entry.New("1", 1),
				entry.New("two", 9),
				entry.New("3", 0),
			},
			putItem: entry.New("", 7),
		},
		{
			name:    "0 items adding a string key",
			initial: []entry.Entry[string, int]{},
			putItem: entry.New("hello", -1000),
		},
		{
			name:    "0 items adding an empty string key",
			initial: []entry.Entry[string, int]{},
			putItem: entry.New("", 9),
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					contains := m.ContainsKey(testCase.putItem.Key())
					assert.False(t, contains)

					prevSize := m.Size()
					m.Put(testCase.putItem.Key(), testCase.putItem.Value())
					assert.Equal(t, prevSize+1, m.Size())

					contains = m.ContainsKey(testCase.putItem.Key())
					assert.True(t, contains)

					value, ok := m.Get(testCase.putItem.Key())
					assert.True(t, ok)
					assert.Equal(t, testCase.putItem.Value(), value)

					prevSize = m.Size()
					newValue := 847348962737
					m.Put(testCase.putItem.Key(), newValue)
					assert.Equal(t, prevSize, m.Size())

					contains = m.ContainsKey(testCase.putItem.Key())
					assert.True(t, contains)

					value, ok = m.Get(testCase.putItem.Key())
					assert.True(t, ok)
					assert.Equal(t, newValue, value)
				})
			}
		})
	}
}

func TestPutExistingKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []entry.Entry[string, int]
		putItem entry.Entry[string, int]
	}{
		{
			name: "3 items adding a new value",
			initial: []entry.Entry[string, int]{
				entry.New("1", 1),
				entry.New("two", 9),
				entry.New("3", 0),
			},
			putItem: entry.New("two", 3),
		},
		{
			name: "3 items adding the same value",
			initial: []entry.Entry[string, int]{
				entry.New("1", 1),
				entry.New("two", 9),
				entry.New("3", 0),
			},
			putItem: entry.New("two", 9),
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					contains := m.ContainsKey(testCase.putItem.Key())
					assert.True(t, contains)

					prevSize := m.Size()
					m.Put(testCase.putItem.Key(), testCase.putItem.Value())
					assert.Equal(t, prevSize, m.Size())

					contains = m.ContainsKey(testCase.putItem.Key())
					assert.True(t, contains)

					value, ok := m.Get(testCase.putItem.Key())
					assert.True(t, ok)
					assert.Equal(t, testCase.putItem.Value(), value)

					prevSize = m.Size()
					newValue := 847348962737
					m.Put(testCase.putItem.Key(), newValue)
					assert.Equal(t, prevSize, m.Size())

					contains = m.ContainsKey(testCase.putItem.Key())
					assert.True(t, contains)

					value, ok = m.Get(testCase.putItem.Key())
					assert.True(t, ok)
					assert.Equal(t, newValue, value)
				})
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		initial      []entry.Entry[int, string]
		putItems     []entry.Entry[int, string]
		expectedSize int
	}{
		{
			name: "adding no items",
			initial: []entry.Entry[int, string]{
				entry.New(9, "foo"),
				entry.New(8, "bar"),
				entry.New(7, "baz"),
			},
			putItems:     []entry.Entry[int, string]{},
			expectedSize: 3,
		},
		{
			name: "adding one item",
			initial: []entry.Entry[int, string]{
				entry.New(9, "foo"),
				entry.New(8, "bar"),
				entry.New(7, "baz"),
			},
			putItems:     []entry.Entry[int, string]{entry.New(1, "hello")},
			expectedSize: 4,
		},
		{
			name: "adding some items that don't exist",
			initial: []entry.Entry[int, string]{
				entry.New(9, "foo"),
				entry.New(8, "bar"),
				entry.New(7, "baz"),
			},
			putItems: []entry.Entry[int, string]{
				entry.New(500, "hello"),
				entry.New(1000, "goodbye"),
			},
			expectedSize: 5,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					m.PutAll(testCase.putItems...)
					assert.Equal(t, testCase.expectedSize, m.Size())

					initialKeys := make([]int, 0)
					for _, entry := range testCase.initial {
						assert.True(t, m.ContainsKey(entry.Key()))
						initialKeys = append(initialKeys, entry.Key())
					}

					putItemKeys := make([]int, 0)
					for _, entry := range testCase.putItems {
						assert.True(t, m.ContainsKey(entry.Key()))
						putItemKeys = append(putItemKeys, entry.Key())
					}

					assert.Len(t, initialKeys, len(testCase.initial))
					assert.True(t, m.ContainsAllKeys(initialKeys...))

					assert.Len(t, putItemKeys, len(testCase.putItems))
					assert.True(t, m.ContainsAllKeys(putItemKeys...))
				})
			}
		})
	}
}

func TestRemoveKeyUntilEmpty(t *testing.T) {
	t.Parallel()

	vals := []entry.Entry[string, int]{
		entry.New("me", 1),
		entry.New("you", 2),
		entry.New("them", 3),
	}

	maps := getMapsForTest(vals...)
	for i := range maps {
		m := maps[i]
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 3, m.Size())
			assert.False(t, m.Empty())
			assert.True(t, m.ContainsKey("me"))
			assert.True(t, m.ContainsKey("you"))
			assert.True(t, m.ContainsKey("them"))

			ok := m.RemoveKey("me")
			assert.True(t, ok)
			assert.False(t, m.ContainsKey("me"))
			assert.Equal(t, 2, m.Size())

			ok = m.RemoveKey("me")
			assert.False(t, ok)
			assert.False(t, m.ContainsKey("me"))

			assert.Equal(t, 2, m.Size())
			assert.True(t, m.ContainsKey("you"))

			ok = m.RemoveKey("you")
			assert.True(t, ok)
			assert.False(t, m.ContainsKey("you"))
			assert.Equal(t, 1, m.Size())

			ok = m.RemoveKey("you")
			assert.False(t, ok)
			assert.False(t, m.ContainsKey("you"))

			assert.Equal(t, 1, m.Size())
			assert.True(t, m.ContainsKey("them"))

			ok = m.RemoveKey("them")
			assert.True(t, ok)
			assert.False(t, m.ContainsKey("them"))
			assert.Equal(t, 0, m.Size())

			ok = m.RemoveKey("them")
			assert.False(t, ok)
			assert.False(t, m.ContainsKey("them"))

			assert.Equal(t, 0, m.Size())
			assert.True(t, m.Empty())
		})
	}
}

func TestRemoveAllKeys(t *testing.T) {
	t.Parallel()

	vals := []entry.Entry[string, int]{
		entry.New("me", 1),
		entry.New("you", 2),
		entry.New("them", 3),
	}

	maps := getMapsForTest(vals...)
	for i := range maps {
		m := maps[i]
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, 3, m.Size())
			assert.True(t, m.ContainsAllKeys("me", "you", "them"))

			removed := m.RemoveAllKeys("me", "you", "them", "we", "us")
			assert.Equal(t, 0, m.Size())
			assert.Equal(t, 3, removed)

			removed = m.RemoveAllKeys("me", "you", "them", "we", "us")
			assert.Equal(t, 0, m.Size())
			assert.Equal(t, 0, removed)

			m.Put("person", -77)
			assert.Equal(t, 1, m.Size())

			removed = m.RemoveAllKeys()
			assert.Equal(t, 1, m.Size())
			assert.Equal(t, 0, removed)

			removed = m.RemoveAllKeys("person")
			assert.Equal(t, 0, m.Size())
			assert.Equal(t, 1, removed)

			m.PutAll(entry.New("me", 20), entry.New("you", 10))
			assert.Equal(t, 2, m.Size())

			removed = m.RemoveAllKeys("me", "them")
			assert.Equal(t, 1, m.Size())
			assert.Equal(t, 1, removed)

			removed = m.RemoveAllKeys("me", "you", "them")
			assert.Equal(t, 0, m.Size())
			assert.Equal(t, 1, removed)
		})
	}
}

func TestAnyAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		initial     []entry.Entry[string, int]
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			initial:     []entry.Entry[string, int]{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no negative values with 1 item",
			initial:     []entry.Entry[string, int]{entry.New("foo", 12)},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name: "negative at index 0",
			initial: []entry.Entry[string, int]{
				entry.New("x", -100),
				entry.New("y", 300),
				entry.New("z", 57),
			},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "negative at index 1",
			initial: []entry.Entry[string, int]{
				entry.New("x", 100),
				entry.New("y", -300),
				entry.New("z", 57),
			},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "negative at index 2",
			initial: []entry.Entry[string, int]{
				entry.New("x", 100),
				entry.New("y", 300),
				entry.New("z", -57),
			},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "no negative values with 3 items",
			initial: []entry.Entry[string, int]{
				entry.New("x", 100),
				entry.New("y", 300),
				entry.New("z", 57),
			},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name: "all negatives",
			initial: []entry.Entry[string, int]{
				entry.New("x", -100),
				entry.New("y", -300),
				entry.New("z", -57),
			},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				m := maps[i]
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					isValueNegative := func(_ string, value int) bool {
						return value < 0
					}
					t.Run("Any", func(t *testing.T) {
						t.Parallel()

						assert.Equal(t, testCase.expectedAny, m.Any(isValueNegative))
					})
					t.Run("All", func(t *testing.T) {
						t.Parallel()

						assert.Equal(t, testCase.expectedAll, m.All(isValueNegative))
					})
				})
			}
		})
	}
}

func TestContainsAnyKeyContainsAllKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		initial     []entry.Entry[string, int]
		keys        []string
		expectedAny bool
		expectedAll bool
	}{
		{
			name:        "no values",
			initial:     []entry.Entry[string, int]{},
			keys:        []string{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name:        "no values with attempted match",
			initial:     []entry.Entry[string, int]{},
			keys:        []string{"foo", "bar", "baz"},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name:        "no values match with 1 item",
			initial:     []entry.Entry[string, int]{entry.New("twelve", 12)},
			keys:        []string{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name: "no values with 3 items",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", 100),
				entry.New("negative", -300),
				entry.New("fiftyseven", 57),
			},
			keys:        []string{},
			expectedAny: false,
			expectedAll: true,
		},
		{
			name: "one match with 3 items",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", 100),
				entry.New("threehundred", 300),
				entry.New("fiftyseven", -57),
			},
			keys:        []string{"f", "a", "threeh", "fiftyseven"},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "no matches values with 3 items",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", 100),
				entry.New("threehundred", 300),
				entry.New("fiftyseven", 57),
			},
			keys:        []string{"fivehundred", "sevenhundred"},
			expectedAny: false,
			expectedAll: false,
		},
		{
			name: "1 of 3 values match",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", -100),
				entry.New("threehundred", -300),
				entry.New("fiftyseven", -57),
			},
			keys:        []string{"threehundred", "fivehundred"},
			expectedAny: true,
			expectedAll: false,
		},
		{
			name: "1 of 3 values all match",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", -100),
				entry.New("threehundred", -300),
				entry.New("fiftyseven", -57),
			},
			keys:        []string{"hundred"},
			expectedAny: true,
			expectedAll: true,
		},
		{
			name: "2 of 3 values all match",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", -100),
				entry.New("threehundred", -300),
				entry.New("fiftyseven", -57),
			},
			keys:        []string{"hundred", "threehundred"},
			expectedAny: true,
			expectedAll: true,
		},
		{
			name: "all values match",
			initial: []entry.Entry[string, int]{
				entry.New("hundred", -100),
				entry.New("threehundred", -300),
				entry.New("fiftyseven", -57),
			},
			keys:        []string{"hundred", "threehundred", "fiftyseven"},
			expectedAny: true,
			expectedAll: true,
		},
	}

	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.initial...)
			for i := range maps {
				mp := maps[i]
				t.Run(fmt.Sprintf("%T", mp), func(t *testing.T) {
					t.Parallel()

					t.Run("ContainsAny", func(t *testing.T) {
						t.Parallel()

						assert.Equal(t, testCase.expectedAny, mp.ContainsAnyKey(testCase.keys...))
					})
					t.Run("ContainsAll", func(t *testing.T) {
						t.Parallel()

						assert.Equal(t, testCase.expectedAll, mp.ContainsAllKeys(testCase.keys...))
					})
				})
			}
		})
	}
}

func TestForEach(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		items    []entry.Entry[string, int]
		expected int
	}{
		{
			name:     "sum nothing",
			items:    []entry.Entry[string, int]{},
			expected: 0,
		},
		{
			name:     "sum a single number",
			items:    []entry.Entry[string, int]{entry.New("foo", 12)},
			expected: 15,
		},
		{
			name: "sum a few numbers",
			items: []entry.Entry[string, int]{
				entry.New("", -100),
				entry.New("b", 300),
				entry.New("cde", 57),
			},
			expected: 261,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.items...)
			for i := range maps {
				m := maps[i]
				total := 0
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					m.ForEach(func(key string, value int) {
						total += len(key) + value
					})
					assert.Equal(t, testCase.expected, total)
				})
			}
		})
	}
}

func TestFindOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		items         []entry.Entry[int, int]
		expectedKey   int
		expectedValue int
	}{
		{
			name: "negative at index (10, -100)",
			items: []entry.Entry[int, int]{
				entry.New(10, -100),
				entry.New(5, 300),
				entry.New(1, 57),
			},
			expectedKey:   10,
			expectedValue: -100,
		},
		{
			name: "negative at key (5, -300)",
			items: []entry.Entry[int, int]{
				entry.New(10, 100),
				entry.New(5, -300),
				entry.New(1, 57),
			},
			expectedKey:   5,
			expectedValue: -300,
		},
		{
			name: "negative at index (1, -57)",
			items: []entry.Entry[int, int]{
				entry.New(10, 100),
				entry.New(5, 300),
				entry.New(1, -57),
			},
			expectedKey:   1,
			expectedValue: -57,
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.items...)
			for i := range maps {
				m := maps[i]
				isKeyPlusValueNegative := func(key int, value int) bool {
					return key+value < 0
				}
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					key, val, ok := m.Find(isKeyPlusValueNegative)
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedKey, key)
					assert.Equal(t, testCase.expectedValue, val)

					valFromMap, ok := m.Get(key)
					assert.True(t, ok)
					assert.Equal(t, testCase.expectedValue, valFromMap)
				})
			}
		})
	}
}

func TestFindNotOk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []entry.Entry[int, int]
	}{
		{
			name:  "no values",
			items: []entry.Entry[int, int]{},
		},
		{
			name: "no negative pair sums with 1 item",
			items: []entry.Entry[int, int]{
				entry.New(1, 12),
			},
		},
		{
			name: "no negative pair sums with 3 items",
			items: []entry.Entry[int, int]{
				entry.New(10, 100),
				entry.New(501, -500),
				entry.New(1, 57),
			},
		},
	}
	for i := range tests {
		testCase := tests[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			maps := getMapsForTest(testCase.items...)
			for i := range maps {
				m := maps[i]
				isKeyPlusValueNegative := func(key int, value int) bool {
					return key+value < 0
				}
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					t.Parallel()

					_, _, ok := m.Find(isKeyPlusValueNegative)
					assert.False(t, ok)
				})
			}
		})
	}
}
