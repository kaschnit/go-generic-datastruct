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
