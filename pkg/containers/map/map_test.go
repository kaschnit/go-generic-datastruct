package mapp_test

import (
	"fmt"
	"testing"

	mapp "github.com/kaschnit/go-ds/pkg/containers/map"
	"github.com/kaschnit/go-ds/pkg/containers/map/entry"
	"github.com/kaschnit/go-ds/pkg/containers/map/hashmap"
	"github.com/stretchr/testify/assert"
)

func getMapsForTest[K comparable, V any](entries ...entry.Entry[K, V]) []mapp.Map[K, V] {
	return []mapp.Map[K, V]{
		hashmap.New(entries...),
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			maps := getMapsForTest(testCase.initial...)
			for _, m := range maps {
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
					assert.False(t, m.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	t.Parallel()

	maps := getMapsForTest[string, int]()
	for _, m := range maps {
		assert.True(t, m.Empty())

		m.Put("foo", 100)
		assert.False(t, m.Empty())

		m.RemoveKey("foo")
		assert.True(t, m.Empty())
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
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			maps := getMapsForTest(testCase.initial...)
			for _, m := range maps {
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
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
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			maps := getMapsForTest(testCase.initial...)
			for _, m := range maps {
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
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
	for _, m := range maps {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			maps := getMapsForTest(testCase.initial...)
			for _, m := range maps {
				t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
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
