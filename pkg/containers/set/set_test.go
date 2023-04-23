package set_test

import (
	"fmt"
	"testing"

	"github.com/kaschnit/go-ds/pkg/containers/set"
	"github.com/kaschnit/go-ds/pkg/containers/set/hashset"
	"github.com/stretchr/testify/assert"
)

func getSetsForTest[T comparable](values ...T) []set.Set[T] {
	return []set.Set[T]{
		hashset.New(values...),
	}
}

func TestEmptyFalse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []string
	}{
		{
			name:    "3 items",
			initial: []string{"a", "b", "c"},
		},
		{
			name:    "1 item",
			initial: []string{"hello"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())
				})
			}
		})
	}
}

func TestEmptyTrue(t *testing.T) {
	t.Parallel()

	sets := getSetsForTest[int]()
	for _, s := range sets {
		assert.True(t, s.Empty())

		s.Add(1)
		assert.False(t, s.Empty())

		s.Remove(1)
		assert.True(t, s.Empty())
	}
}

func TestSize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  []float64
		expected int
	}{
		{
			name:     "3 items",
			initial:  []float64{1.2, 2.3, 999.999},
			expected: 3,
		},
		{
			name:     "1 item",
			initial:  []float64{7.000},
			expected: 1,
		},
		{
			name:     "0 items",
			initial:  []float64{},
			expected: 0,
		},
		{
			name:     "6 items",
			initial:  []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
			expected: 6,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, q := range sets {
				t.Run(fmt.Sprintf("%T", q), func(t *testing.T) {
					assert.Equal(t, testCase.expected, q.Size())
					assert.Equal(t, len(testCase.initial), q.Size())
				})
			}
		})
	}
}

func TestClearNonEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []float64
	}{
		{
			name:    "3 items",
			initial: []float64{1.2, 2.3, 999.999},
		},
		{
			name:    "1 item",
			initial: []float64{7.000},
		},
		{
			name:    "6 items",
			initial: []float64{2.5, 1.000, -5.444, 0.1, 500, 12},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())

					s.Add(1.2345)
					assert.False(t, s.Empty())

					s.Clear()
					assert.True(t, s.Empty())
				})
			}
		})
	}
}

func TestClearEmpty(t *testing.T) {
	t.Parallel()

	sets := getSetsForTest[string]()
	for _, s := range sets {
		assert.True(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())

		s.Add("hello")
		assert.False(t, s.Empty())

		s.Clear()
		assert.True(t, s.Empty())
	}
}

func TestAddNewItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		initial  []string
		pushItem string
	}{
		{
			name:     "3 items pushing a string",
			initial:  []string{"1", "two", "3"},
			pushItem: "foo",
		},
		{
			name:     "3 items pushing an empty string",
			initial:  []string{"1", "two", "3"},
			pushItem: "",
		},
		{
			name:     "0 items pushing a string",
			initial:  []string{},
			pushItem: "hello",
		},
		{
			name:     "0 items pushing an empty string",
			initial:  []string{},
			pushItem: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					contains := s.Contains(testCase.pushItem)
					assert.False(t, contains)

					prevSize := s.Size()
					s.Add(testCase.pushItem)
					assert.Equal(t, prevSize+1, s.Size())

					contains = s.Contains(testCase.pushItem)
					assert.True(t, contains)
				})
			}
		})
	}
}

func TestAddExistingItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial []string
		addItem string
	}{
		{
			name:    "1 items pushing a string",
			initial: []string{"abc"},
			addItem: "abc",
		},
		{
			name:    "3 items pushing a string",
			initial: []string{"1", "two", "3"},
			addItem: "3",
		},
		{
			name:    "3 items pushing an empty string",
			initial: []string{"1", "", "3"},
			addItem: "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					contains := s.Contains(testCase.addItem)
					assert.True(t, contains)

					prevSize := s.Size()
					s.Add(testCase.addItem)
					assert.Equal(t, prevSize, s.Size())

					contains = s.Contains(testCase.addItem)
					assert.True(t, contains)
				})
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		initial      []int
		addItems     []int
		expectedSize int
	}{
		{
			name:         "adding no items",
			initial:      []int{9, 8, 7},
			addItems:     []int{},
			expectedSize: 3,
		},
		{
			name:         "adding one item",
			initial:      []int{9, 8, 7},
			addItems:     []int{1},
			expectedSize: 4,
		},
		{
			name:         "adding some items that don't exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{500, 1000},
			expectedSize: 5,
		},
		{
			name:         "adding some items that exist and don't exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{9, 1000},
			expectedSize: 4,
		},
		{
			name:         "adding some items that already exist",
			initial:      []int{9, 8, 7},
			addItems:     []int{9, 7},
			expectedSize: 3,
		},
		{
			name:         "adding one item to an empty list",
			initial:      []int{},
			addItems:     []int{-52},
			expectedSize: 1,
		},
		{
			name:         "addding some items to an empty set",
			initial:      []int{},
			addItems:     []int{1, 2, 3},
			expectedSize: 3,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			sets := getSetsForTest(testCase.initial...)
			for _, s := range sets {
				t.Run(fmt.Sprintf("%T", s), func(t *testing.T) {
					s.AddAll(testCase.addItems...)
					assert.Equal(t, testCase.expectedSize, s.Size())
					assert.True(t, s.ContainsAll(testCase.initial...))
					assert.True(t, s.ContainsAll(testCase.addItems...))
				})
			}
		})
	}
}

func TestRemoveUntilEmpty(t *testing.T) {
	t.Parallel()

	vals := []int{1, 2, 3}
	sets := getSetsForTest(vals...)
	for _, s := range sets {
		assert.Equal(t, 3, s.Size())
		assert.True(t, s.Contains(1))

		ok := s.Remove(1)
		assert.True(t, ok)
		assert.False(t, s.Contains(1))

		ok = s.Remove(1)
		assert.False(t, ok)
		assert.False(t, s.Contains(1))

		assert.Equal(t, 2, s.Size())
		assert.True(t, s.Contains(2))

		ok = s.Remove(2)
		assert.True(t, ok)
		assert.False(t, s.Contains(2))

		ok = s.Remove(2)
		assert.False(t, ok)
		assert.False(t, s.Contains(2))

		assert.Equal(t, 1, s.Size())
		assert.True(t, s.Contains(3))

		ok = s.Remove(3)
		assert.True(t, ok)
		assert.False(t, s.Contains(3))

		ok = s.Remove(3)
		assert.False(t, ok)
		assert.False(t, s.Contains(3))

		assert.Equal(t, 0, s.Size())
	}
}
